# collect

Lazy, chainable collections for Go 1.27, a port of Laravel's LazyCollection. See README.md for the API, and its tables for how each Laravel method maps.

## Rules

### Laziness

- Every operation is lazy unless it is a terminal one (`All`, `Count`, `First`, `SumBy` and the like). A lazy method returns a new value whose sequence runs only when read.
- A method that must read every value first, such as `SortBy` or `GroupBy`, does so when its result is read, not when it is called. Only `Eager`, `Ensure`, `Combine` and `RandomN` read when called, as their names or results say.
- A read of a result reads the source at most once. A method with two results reads once for both and keeps them, as `Partition` does. A method that repeats keeps what it read, as `Multiply` does. Laravel reads them once too.
- `IsEmpty`, `IsNotEmpty`, `HasSole` and `HasMany` read values only to answer, as Laravel's do. Their docs say that a source that can be read once loses them, and to call `Remember` first. Keep that note on any method like them.
- Every loop over a source returns as soon as `yield` returns false. Each lazy method's test calls `assertStopsEarly` or `assertStopsEarly2` on an endless source, for every type that has the method.
- `Of` copies its arguments. A chunk, a window or a group gets slices of its own. A value a caller holds never changes when the source does.
- `FromSlice` and `FromMap` do not copy, on purpose: they are the way to read a large slice or map without one. Do not add a copy to them. A source that copies has no reason to exist beside `Of`.

### Types and the instantiation cycle

- A method whose result nests the receiver's type pins a type parameter instead of naming the nested type. `Chunk` is `Chunk[C interface{ Collection[T] }](size int) Collection[C]`, and converts each chunk with `C(...)`. Go infers `C`, so callers never write it. `GroupBy`, `ChunkWhile`, `ChunkBy`, `Sliding`, `Split` and `SplitIn` do the same, and `Zip`, `CrossJoin` and `Entries` pin `E interface{ Entry[...] }`.
- Never name a nested type in a method, in its signature or its body: `Collection[Collection[T]]`, or `Collection[Entry[K, V]]` in a method of `KeyedCollection`. It is an "instantiation cycle" (golang/go#80172): the compiler instantiates every method of the nested type, then of a deeper one, without end. A pinned `C` is instantiated only by a call, so only as deep as calls go. Package functions and the caller's code are not checked this way.
- A method must not call a pinned method of the package's own types either, such as `c.Chunk(n)` inside `SplitIn`. The call instantiates `C` as the nested type, which is the same cycle. Loop by hand. To keep pairs, use `pairsOf` and `keyedOf` in `collection.go`, or a `rememberState[Pair[K, V]]`, which has no methods that lead back.
- `Collection[T any]` cannot add a constraint to `T` in a method. A method that needs numbers, an order or equality takes a key function with its own type parameter, such as `SumBy[N Number](f func(T) N)`. Callers pass `Identity` for Laravel's form with no key. Do not use `any` as a map key to get equality: it compiles for every `T` and panics on a slice.

### Both types

- A method exists on both `Collection` and `KeyedCollection`. On `KeyedCollection`, its functions take `(K, V)`, it keeps each pair's key, and a result with an item returns `(K, V, bool)` or `(K, V, error)`. `Value` and `Key` are the `Identity` of such a function.
- The exception: a function that a method also applies to another collection, or to one side of a pair only, takes only what it compares. `DiffBy`, `IntersectBy`, `DiffAssocBy`, `IntersectAssocBy`, `DiffFunc`, `IntersectFunc` and `FlipBy` take values, and `SortKeysBy` and `SortKeysFunc` keys, so `Identity` fits them.
- A `KeyedCollection` keeps every key, even where Laravel numbers the pairs again (`Nth`, `Multiply`, `Shuffle`, `RandomN`, `Split`, `FlatMap`). The README lists this difference.
- These exist on one type only, by design. Keep it so, and add a new exception here with its reason.

| Only on | Methods | Why |
| --- | --- | --- |
| `Collection` | `Keyed` | A `KeyedCollection` has keys already |
| `Collection` | `All` | A `KeyedCollection` has `ToMap`, whose name says a map has no order, and `Entries().All()` in order |
| `Collection` | `Pad` | A value it adds has no key |
| `Collection` | `Zip`, `CrossJoin` | A pair has no key, and Laravel drops the keys too. Call `Values()` first |
| `KeyedCollection` | `Keys`, `Values`, `Get`, `MustGet`, `Has`, `HasAny`, `Only`, `Except`, `FlipBy`, `Union`, `Replace`, `DiffKeys`, `IntersectKeys`, `DiffAssocBy`, `IntersectAssocBy`, `SortKeysBy`, `SortKeysByDesc`, `SortKeysFunc`, `ToMap` | They work on keys. Call `Keyed()` first |

### Results and failures

| A method returns | When |
| --- | --- |
| `(T, bool)` | There may be nothing, and that is the only way to fail: `First`, `Get`, `MaxBy` |
| `(T, error)` | It can fail in more than one way, or the reason matters: `Sole`, `Ensure`, `Combine`, `RandomN` |
| `T`, and panics | `Must` plus the name of the method above, for a caller who knows it cannot fail. It panics with that method's error |
| A panic, with no `Must` | An int argument that cannot be right, where Laravel throws too: `Chunk(0)`, `Sliding(0, 1)`, `Nth(0, 0)`, `Split(0)`, `SplitIn(0)`, `RangeStep` with step 0. It panics with an error that wraps `ErrInvalidArgument`, through `invalidArgument` |

- Every error is, or wraps, one of the package's `Err*` values, with the details in its message.
- A Laravel method that throws maps to a `Must` method: `firstOrFail` is `MustFirst`, `sole` is `MustSole`, `ensure` is `MustEnsure`.
- Otherwise follow Laravel's meaning, even where Go would differ: `Take(-n)` takes the last values, `Slice` and `ForPage` take array_slice's negative values, and `RandomN(-1)` is empty. `Skip(-n)` is the exception: Laravel's lazy and eager forms disagree, so a negative count is 0.
- An int argument may be anything from `math.MinInt` to `math.MaxInt`. Do not allocate by it before the values come (`Chunk` caps its first buffer), and do not let arithmetic on it overflow (`-math.MinInt`, `(page-1)*perPage`).
- A map keyed by what a key function returns finds a group by its position, through `groupsOf` in `collection.go`. Looking the key up again loses a NaN, which never equals itself.
- A method that keeps state across reads (`Partition`, `Remember`) keeps nothing from a read that panicked, or keeps the panic. A later read must never return part of the values as if they were all.

### Names

- A Laravel alias is not ported: `average`, `some`, `containsOneItem` and the like. List it in the README's table of other names instead.
- The exception is a name Go users expect: `JoinBy` is `ImplodeBy`, under the name of `strings.Join`. Laravel's `join` with a final glue is `JoinFinalBy`.
- A method whose function returns a key, to compare, order, group or add up, ends in `By`: `SumBy`, `UniqueBy`, `SortBy`, `ImplodeBy`, `WhereInBy`. It is the method that `Identity` turns into Laravel's form with no key. A method whose function returns a bool or a new value has no suffix: `Filter`, `Map`.
- A method whose function compares two values ends in `Func`, as in the standard library: `sort($callback)` is `SortFunc`, `diffUsing` is `DiffFunc`.
- `By` is only for a method that takes a key function. Laravel's `intersectByKeys` is `IntersectKeys`, as `diffKeys` is `DiffKeys`.
- A method that takes another collection takes it first, as Laravel does: `DiffBy(other, f)`, `DiffFunc(other, eq)`. Otherwise the function comes first: `WhereInBy(f, keys...)`, `ImplodeBy(f, sep)`.
- A method that returns a collection of a `KeyedCollection`'s own pairs pins a type parameter too, as `Entries` does: `Entries[E interface{ Entry[K, V] }]() Collection[E]`.

### JSON

- A collection implements only `MarshalJSONTo` of `encoding/json/v2`. Go 1.27's `encoding/json` calls it too, so there is no `MarshalJSON`.
- A `KeyedCollection` streams its pairs in order, always as an object. A key is encoded with the encoder's options, and a key that is not a JSON string, number or boolean is an error that wraps `ErrUnsupportedKey`. A key that comes twice is an error in `encoding/json/v2`, which rejects duplicate names by default, and two members in `encoding/json`.

### Tests

- Every exported function and method has an example in its `_test.go`. Prefer Laravel's own example for the method. An example's `Output` is the check.
- A `Test` covers what an example would not show: empty sources, negative and zero counts, panics, errors, laziness and early breaks.
- Keep statement coverage at 100%.
- Check every error, or drop it with `_ =` and a comment saying why. errcheck runs with none of its default exclusions. Only `fmt.Println` in `_test.go` files is excluded, for the examples.
- Release with the `Tag and Release` workflow: `gh workflow run tag_and_release.yml -f version=vX.Y.Z`. It tags main's head only once its tests passed, and never reuses a tag: Go caches a version forever.

## Skills

Read the matching one in `.agents/skills/` before you start.

| Skill | When |
| --- | --- |
| `declscope-authoring` | Adding, naming or moving any declaration |
| `declscope-adoption` | Changing declscope's configuration |

## Layout

| File | Holds |
| --- | --- |
| `collection.go` | Core and shared: every type, the sources `Of` and `From*`, `Identity`, `Value`, `Key`, `Number`, `Entry`, the `Err*` values, and helpers several files use |
| `<method>.go` | One method, for every type that has it: `map.go` holds the `Map` of `Collection` and of `KeyedCollection` |
| `<method>_test.go` | The examples and tests of that method, for every type |
| `fixture_test.go` | Test helpers that several test files use, each `//declscope:shared` |

A file is named for its method in snake case (`key_by.go` for `KeyBy`). So the method's name carries the file's declscope namespace, and a private helper of the file starts with it (`chunkCheckSize` in `chunk.go`). Two methods that mirror each other, such as `Filter` and `Reject`, still get a file each. A `Must` method gets its own file too (`must_first.go`).

## Before you finish

```sh
go vet ./... && go test -race ./...
mise exec -- golangci-lint run ./...
mise exec -- declscope shrink ./...
mise exec -- declscope ./...
```

declscope runs with `.declscope.yaml` (`qualify: ondemand`, `exported: true`).

## Rejected designs

| Idea | Why not |
| --- | --- |
| Package functions for `Map` and other type-changing operations | That is what the other ports do, and it breaks the chain. Go 1.27's generic methods are the reason this module exists. |
| A file per type, such as `keyed.go` for all of `KeyedCollection` | Most methods exist on both types. One file per method shows which types have it, and which lack it. |
| `Chunk` returning `Collection[Collection[T]]` | It does not compile: "instantiation cycle" (golang/go#80109, golang/go#80172). The Go team does not plan to relax it. The pinned `C` gives the same type. |
| `Chunks[T]` and `KeyedChunks[K, V]`, types of their own for chunks | They did compile, but each needed its own `Map`, `Count` and the rest. They could not chunk again or take `Keyed`, as both were cycles. With a pinned `C`, a chunk is a `Collection`, and every method applies. |
| A `Flatten` method on `Collection` | A method cannot require `T` to be a collection. `FlatMap` with a function that returns its argument does it. |
| `Unique()` with `any` as the key of a map | It compiles for every `T`, then panics at run time for a slice: "hash of unhashable type". `UniqueBy(Identity)` is checked by the compiler. |
| `FirstOrFail` returning `(T, error)` | Its only error was "not found", which `First`'s `bool` already says. Laravel's `firstOrFail` throws, so it is `MustFirst`. |
| `MustFirstWhere`, `MustMax` and other `Must` methods for every `(T, bool)` | `Filter(f).MustFirst()` does the first. For the rest, the caller knows whether the collection is empty. Add one only when it is asked for. |
| An eager `Collection` beside a lazy one, like Laravel | Each method would be written twice, and no interface can share them: interface methods cannot have type parameters. `All()` reads a lazy one into a slice. |
| Laravel's keys on every collection | Go has no ordered map, and the kept keys are a well-known trap (`values()` after `filter`). Keys are opt-in instead: `Keyed()` and `KeyBy()` return a `KeyedCollection` over `iter.Seq2`, and `Values()` drops them. |
| Unique keys in a `KeyedCollection` | A lazy sequence would have to remember every key it passed. Pairs flow as they come, and `ToMap()` collapses them into a map. |
| `MarshalJSON` of `encoding/json` | Go 1.27's `encoding/json` calls `MarshalJSONTo`, so `MarshalJSON` never ran. `MarshalJSONTo` also streams the values as they are read. |
| A single `From`, for `iter.Seq` only | The keyed collection needs its own source, and `From2` reads badly. `FromSeq` and `FromSeq2` follow `iter.Seq` and `iter.Seq2`. |
| `Unkeyed()` as the name of the way back from `Keyed()` | `Values()` pairs with `Keys()`, as `maps.Keys` and `maps.Values` do, and is Laravel's name. |
| `Sum`, `Unique`, `Implode` and the like, with no `By` | They took a key function as `SortBy` and `GroupBy` do, under another kind of name. Every method that takes a key ends in `By`, as in samber/lo. |
| `WithHeartbeat`, `Throttle`, `TakeUntilTimeout` | They looked at the clock only when a value came, so a slow source delayed every beat and every deadline. `Throttle` slept inside the iterator, which no `context` could stop. |
| `WhenEmpty`, `WhenNotEmpty` | Deciding when read meant reading the first value twice: a source that can be read once lost it. Deciding when called would read the source early. |
| `Partition` returning two `Filter`s, and `Multiply` reading the source `n` times | A source that can be read once gave an empty second half, or a single copy. |
| `IntersectByKeys`, Laravel's name | It has `By` and no key function. It is `IntersectKeys`, beside `DiffKeys`. |
| `ModeBy` returning `[]K` | Every other method that reads every value first returns a collection. It returns `Collection[K]`. |
| `sync.Once` in `Partition` | A panic during the read counted as done, so later reads returned part of the values. A mutex and a flag set only when the read ends keep nothing from a panic. |
| `Remember` taking a source's panic as its end | iter.Pull reports the end after a panic, so later reads returned the kept values only. It keeps the panic, and repeats it. |
| `RangeStep` stopping when `to - n < step` | `to - n` overflows for a signed span past the type's maximum, and a float's last value was lost to rounding. It adds the step, then stops when the next number wraps, does not change, or passes `to`. |
| `KeyedCollection.All` returning `map[K]V` | The name said "all", but a map drops the order and every repeated key: sorting, then `All()`, gave an unsorted map. It is `ToMap`, and `Entries().All()` keeps both. |
| `Pairs`, a package function, and `Pair{First, Second}` | A method can return the pairs lazily by pinning a type parameter, so it is `Entries`, after ECMAScript. `Entry{Key, Value}` names what a pair holds. |
| `Tap` | It called its function with the collection when called. Any look at the whole, such as `Count()`, read the source early. A source that can be read once was then empty for the rest of the chain. `TapEach` sees each value lazily, and `Pipe` calls a function in a chain. |
| `ReduceInto` | It passed the result by pointer, as Laravel passes it by reference, because PHP copies arrays. A Go map or slice shares its contents, so `Reduce` does the same with no `(*m)[k]`. |
| `Pluck("user.name")` and other string paths | They need reflection, and lose the types. Pass a function to `Map`. |

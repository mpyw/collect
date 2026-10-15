# collect

[![Test](https://github.com/mpyw/collect/actions/workflows/test.yml/badge.svg)](https://github.com/mpyw/collect/actions/workflows/test.yml)
[![Coverage](https://codecov.io/gh/mpyw/collect/graph/badge.svg)](https://codecov.io/gh/mpyw/collect)
[![Go Reference](https://pkg.go.dev/badge/github.com/mpyw/collect.svg)](https://pkg.go.dev/github.com/mpyw/collect)

Lazy, chainable collections for Go, in the manner of [Laravel's LazyCollection](https://laravel.com/docs/collections#lazy-collections).

> [!WARNING]
> This is an early draft. The API may change.

```go
names := collect.Of(users...).
	Filter(func(u User) bool { return u.Age >= 20 }).
	Map(func(u User) string { return u.Name }).
	All()
```

collect needs Go 1.27 or later.
Go 1.27 lets a method declare its own type parameters.
So `Map` can turn a `Collection[User]` into a `Collection[string]` in the middle of a chain.

Every method has an example on [pkg.go.dev](https://pkg.go.dev/github.com/mpyw/collect).
Most of them are Laravel's own examples.

## Install

```sh
go get github.com/mpyw/collect
```

## Lazy by default

A `Collection[T]` wraps an `iter.Seq[T]`.
A method builds a new sequence and runs nothing.
The values flow only when a terminal method, such as `All()`, reads them.
So an endless sequence is fine, as long as something ends it:

```go
squares := collect.FromSeq(naturals).
	Map(func(n int) int { return n * n }).
	Take(5).
	All()
// [1 4 9 16 25]
```

| Source | What it reads |
| --- | --- |
| `Of(values...)` | A copy of the values |
| `FromSlice(s)` | `s`, with no copy |
| `FromSeq(seq)` | An `iter.Seq[T]` |
| `FromSeq2(seq)`, `FromMap(m)` | An `iter.Seq2[K, V]` or a map, as a keyed collection |
| `Range(from, to)`, `RangeStep(from, to, step)` | Numbers, counting up or down |
| `Times(n, f)` | What `f` returns for 1 to `n` |
| `Combine(keys, values)` | Two collections, as the keys and the values of a keyed collection |

> [!NOTE]
> A collection is as reusable as its source.
> One from `Of` or `FromSlice` can be read any number of times.
> One from `FromSeq` can be read as many times as its `iter.Seq` allows.
> `Remember()` and `Eager()` make any collection reusable.

`Eager()` reads every value when it is called.
`Remember()` reads each value of the source once, as it is first needed, and keeps it:

```go
users := collect.FromSeq(cursor).Remember()

users.Take(5).All()  // Reads 5 values from the cursor.
users.Take(20).All() // Returns the 5 kept, and reads 15 more.
```

> [!WARNING]
> `Remember()` keeps every value it has read.
> It holds the source open between reads, until the source ends or the collection is garbage collected.
> It is not safe for reads from several goroutines at once.

> [!WARNING]
> `FromSlice` and `FromMap` do not copy.
> A change to the slice or the map shows in the next read, as with `slices.Values`.
> Use `Of(s...)` for a copy.

Some methods must read every value before they yield the first.
They are still lazy: they read when the result is read, not when they are called.

| Reads every value first | Why |
| --- | --- |
| `SortBy`, `SortFunc`, `Reverse`, `Shuffle` | The first value is known only at the end |
| `GroupBy`, `CountBy`, `MapToGroups` | A group is complete only at the end |
| `Partition` | It reads once for both of its results, and keeps them |
| `Take(-n)`, `Slice` with a negative value, `Pad(-n, v)`, `Split`, `SplitIn` | They count from the end |
| `DiffBy`, `IntersectBy` and the like | They read the other collection whole |

## Key functions and `Identity`

A method cannot add a constraint to `T`, so `Collection[T]` has no `Sum()` for numbers only.
Instead, a method that needs numbers, an order or equality takes a key function, and its name ends in `By`.
Pass `collect.Identity` when the values are their own keys:

```go
c := collect.Of(3, 1, 2)

c.SumBy(collect.Identity)            // 6
c.SortBy(collect.Identity).All()     // [1 2 3]
c.UniqueBy(collect.Identity).Count() // 3
```

Go infers the type of `Identity` from the collection.

| A function that | Name | Examples |
| --- | --- | --- |
| Returns a key to compare, order, group or add up | Ends in `By` | `SumBy`, `MaxBy`, `UniqueBy`, `DiffBy`, `SortBy`, `GroupBy`, `ImplodeBy`, `WhereInBy` |
| Returns a bool | No suffix | `Filter`, `Contains`, `Every`, `TakeWhile`, `Partition` |
| Returns a new value | No suffix | `Map`, `FlatMap`, `MapWithKeys`, `Reduce` |
| Compares two values | Ends in `Func`, as in the standard library | `SortFunc`, `DiffFunc`, `IntersectFunc`, `SortKeysFunc` |

## Results: `bool`, `error` and `Must`

| A method returns | When | Examples |
| --- | --- | --- |
| `(T, bool)` | There may be nothing, and that is the only way to fail | `First`, `Last`, `Get`, `Search`, `MaxBy`, `AvgBy` |
| `(T, error)` | It can fail in more than one way, or the reason matters | `Sole`, `Ensure`, `Combine`, `RandomN` |
| `T`, and panics | You know it cannot fail. Its name is `Must` and the name of the method above | `MustFirst`, `MustLast`, `MustGet`, `MustSole`, `MustEnsure`, `MustCombine` |

Every error wraps one of the package's errors, so `errors.Is` tells them apart.
A `Must` method panics with the same error.

```go
_, err := collect.Of(1, 2, 3).Sole()
errors.Is(err, collect.ErrMultipleItemsFound) // true
```

> [!IMPORTANT]
> An int argument that cannot be right panics, with no `Must` in the name.
> `Chunk(0)`, `Sliding(0, 1)`, `Nth(0, 0)`, `Split(0)` and `RangeStep(0, 1, 0)` are such calls, as `slices.Chunk(s, 0)` is.

## Keys when you want them

A `Collection[T]` has no keys.
`Keyed()` gives each value its index as a key, and returns a `KeyedCollection[int, T]`.
`KeyBy(f)` gives each value the key `f` returns.
`Values()` drops the keys again, and `Keys()` keeps only the keys:

```go
odd := collect.Of(10, 11, 12, 13).Keyed().
	Filter(func(i, v int) bool { return v%2 == 1 })

odd.Keys().All()   // [1 3]
odd.Values().All() // [11 13]
```

A `KeyedCollection[K, V]` wraps an `iter.Seq2[K, V]`.
It has every method of `Collection` but `Keyed`, `Pad`, `Zip` and `CrossJoin`, and each keeps the keys.
Its functions take the key first, as `maps.All` yields it.
A method that returns a value returns the pair instead: `First()` is `(K, V, bool)`.
It also has the methods of keys: `Get`, `Has`, `Only`, `Except`, `FlipBy`, `Union`, `Replace`, `DiffKeys`, `SortKeysBy` and more.

> [!NOTE]
> A key may come more than once, as in an `iter.Seq2`.
> `Count()` and `Keys()` see each one, and `Get` returns the first value of a key.
> `All()` returns a `map[K]V`, so the last value of a key wins.

## Nested collections

`Chunk` returns a `Collection` of collections.
It is a `Collection` like any other, so every method applies to it, `Chunk` included:

```go
quads := collect.Of(1, 2, 3, 4, 5, 6, 7, 8).Chunk(2).Chunk(2)
// [[[1 2] [3 4]] [[5 6] [7 8]]]
```

`GroupBy`, `ChunkWhile`, `ChunkBy`, `Sliding`, `Split` and `SplitIn` nest the same way.
`FlatMap(collect.Identity)` flattens chunks again.

> [!NOTE]
> In the documentation, `Chunk` reads `Chunk[C interface{ Collection[T] }](size int) Collection[C]`.
> `C` is always `Collection[T]`, and Go infers it. Write `c.Chunk(2)`.
>
> A method of `Collection[T]` cannot return `Collection[Collection[T]]` in Go.
> The compiler reports an "instantiation cycle" ([golang/go#80172](https://github.com/golang/go/issues/80172)).
> The type parameter `C` avoids that, and the result is the same type.
> `Zip` and `CrossJoin` return `Pair`s the same way.

## JSON

A collection streams itself as JSON as it is read, through `encoding/json/v2`'s `MarshalJSONTo`.
`json.Marshal` of `encoding/json` uses it too.
A `Collection` is an array, and a `KeyedCollection` is an object in its own order:

```go
b, _ := json.Marshal(collect.Of("x", "y").Keyed())
// {"0":"x","1":"y"}
```

## Differences from Laravel

| Laravel | collect | Why |
| --- | --- | --- |
| Keys survive `filter` | Only after `Keyed()` | You choose a `Collection` or a `KeyedCollection` |
| `fn($value, $key)` | `func(key, value)` | The order of `iter.Seq2` |
| `sum()`, `sort()`, `unique()` | `SumBy(collect.Identity)` and the like | See [Key functions and `Identity`](#key-functions-and-identity) |
| `first()` returns `null` | `First()` returns `(T, bool)` | Go has no `null` for every type |
| `firstOrFail()`, `sole()` throw | `MustFirst()`, `MustSole()` panic. `Sole()` returns an error | See [Results](#results-bool-error-and-must) |
| `contains($value)`, `where('age', '>', 20)` | `Contains(func...)`, `Filter(func...)` | The compiler checks a function |
| `pluck('user.name')`, `$c->map->name` | `Map(func...)` | No reflection, and no magic properties |
| `*Strict` methods | The methods without `Strict` | Go always compares strictly |
| `join(', ', ' and ')` | `JoinFinalBy(f, ", ", " and ")` | `JoinBy` is `ImplodeBy`, under the name of `strings.Join` |
| `collect($map)` keeps the order | `FromMap(m)` has no order | A Go map has no order. Sort it, or use `FromSeq2` |
| `chunk(0)` is empty | `Chunk(0)` panics | The same as `slices.Chunk` |
| `zip` takes many arrays | `Zip` takes one collection | Go has no variadic type parameters |

Some Laravel methods have another name or another way here:

| Laravel | collect |
| --- | --- |
| `average`, `some` | `AvgBy`, `Contains` |
| `hasSole()`, `containsOneItem()` | The error of `Sole()`, which returns the value too |
| `hasMany()`, `containsManyItems()` | `c.Take(2).Count() == 2` |
| `merge` | `Concat` for a `Collection`, `Replace` for a `KeyedCollection` |
| `mapToDictionary` | `MapToGroups` |
| `reduceWithKeys` | `KeyedCollection.Reduce` |
| `sort($callback)`, `sortKeysUsing`, `diffUsing`, `intersectUsing` | `SortFunc`, `SortKeysFunc`, `DiffFunc`, `IntersectFunc` |
| `value('name')`, `firstWhere('name', 'x')` | `Map(f).First()`, `FirstWhere(func...)` |
| `isEmpty()`, `isNotEmpty()` | The `bool` of `First()`: `_, ok := c.First()` |
| `collapse()`, `flatten(1)` | `FlatMap(collect.FromSlice)`, `FlatMap(collect.Identity)` |
| `toArray()`, `getIterator()`, `toJson()`, `toPrettyJson()` | `All()`, `Seq()`, `json.Marshal(c)`, `json.Marshal(c, jsontext.WithIndent("  "))` |
| `make`, `wrap`, `empty`, `fromJson` | `Of`, `FromSlice`, the zero value, `json.Unmarshal` then `FromSlice` |

These have no Go equivalent, and are not here:

| Laravel | Why |
| --- | --- |
| `dd`, `dump` | Use `fmt` |
| `macro`, `proxy`, `mapInto`, `pipeInto` | PHP's classes and magic methods |
| `eachSpread`, `mapSpread`, `reduceSpread` | Go has no argument spread |
| `flatten()` of any depth, `dot`, `undot`, `mergeRecursive`, `replaceRecursive` | A Go type has a fixed depth |
| `whereNull`, `whereNotNull`, `filter()` with no callback | Go has no `null` or falsy value for every type. Pass a function to `Filter` |
| `diffAssocUsing`, `diffKeysUsing`, `intersectAssocUsing` | Go keys are comparable |
| `select`, `escapeWhenCastingToString`, `getCachingIterator` | PHP's arrays and strings |
| `withHeartbeat`, `throttle`, `takeUntilTimeout` | They look at the clock only when a value comes, so a slow source stops nothing. A Go program times work with `context`, `time.Ticker` or `golang.org/x/time/rate` |
| `whenEmpty`, `whenNotEmpty`, `unlessEmpty`, `unlessNotEmpty` | They read the first value twice, which breaks a source that can be read once. Use `First()`, which returns the value it reads |
| `tap` | To see the whole collection, it must read the source, which may then be empty for the rest of the chain. `TapEach` sees each value lazily. Use `Pipe` to call a function in a chain |
| `reduceInto` | PHP copies an array that a callback returns, so Laravel passes the result by reference. A Go map or slice shares its contents, so `Reduce` does the same |

## License

MIT

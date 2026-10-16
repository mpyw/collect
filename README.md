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
So [`Map`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map) can turn a [`Collection[User]`](https://pkg.go.dev/github.com/mpyw/collect#Collection) into a [`Collection[string]`](https://pkg.go.dev/github.com/mpyw/collect#Collection) in the middle of a chain.

Every method has an example on [pkg.go.dev](https://pkg.go.dev/github.com/mpyw/collect).
Most of them are Laravel's own examples.

## Install

```sh
go get github.com/mpyw/collect
```

## Lazy by default

A [`Collection[T]`](https://pkg.go.dev/github.com/mpyw/collect#Collection) wraps an [`iter.Seq[T]`](https://pkg.go.dev/iter#Seq).
A method builds a new sequence and runs nothing.
The values flow only when a terminal method, such as [`All()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.All), reads them.
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
| [`Of(values...)`](https://pkg.go.dev/github.com/mpyw/collect#Of) | A copy of the values |
| [`FromSlice(s)`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice) | `s`, with no copy |
| [`FromSeq(seq)`](https://pkg.go.dev/github.com/mpyw/collect#FromSeq) | An [`iter.Seq[T]`](https://pkg.go.dev/iter#Seq) |
| [`FromSeq2(seq)`](https://pkg.go.dev/github.com/mpyw/collect#FromSeq2), [`FromMap(m)`](https://pkg.go.dev/github.com/mpyw/collect#FromMap) | An [`iter.Seq2[K, V]`](https://pkg.go.dev/iter#Seq2) or a map, as a keyed collection |
| [`Range(from, to)`](https://pkg.go.dev/github.com/mpyw/collect#Range), [`RangeStep(from, to, step)`](https://pkg.go.dev/github.com/mpyw/collect#RangeStep) | Numbers, counting up or down |
| [`Times(n, f)`](https://pkg.go.dev/github.com/mpyw/collect#Times) | What `f` returns for 1 to `n` |
| [`Combine(keys, values)`](https://pkg.go.dev/github.com/mpyw/collect#Combine) | Two collections, as the keys and the values of a keyed collection |

> [!NOTE]
> A collection is as reusable as its source.
> One from [`Of`](https://pkg.go.dev/github.com/mpyw/collect#Of) or [`FromSlice`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice) can be read any number of times.
> One from [`FromSeq`](https://pkg.go.dev/github.com/mpyw/collect#FromSeq) can be read as many times as its [`iter.Seq`](https://pkg.go.dev/iter#Seq) allows.
> [`Remember()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Remember) and [`Eager()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Eager) make any collection reusable.

[`Eager()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Eager) reads every value when it is called.
[`Remember()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Remember) reads each value of the source once, as it is first needed, and keeps it:

```go
users := collect.FromSeq(cursor).Remember()

users.Take(5).All()  // Reads 5 values from the cursor.
users.Take(20).All() // Returns the 5 kept, and reads 15 more.
```

> [!WARNING]
> [`Remember()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Remember) keeps every value it has read.
> It holds the source open between reads, until the source ends or the collection is garbage collected.
> It is not safe for reads from several goroutines at once.

> [!WARNING]
> [`FromSlice`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice) and [`FromMap`](https://pkg.go.dev/github.com/mpyw/collect#FromMap) do not copy.
> A change to the slice or the map shows in the next read, as with [`slices.Values`](https://pkg.go.dev/slices#Values).
> Use [`Of(s...)`](https://pkg.go.dev/github.com/mpyw/collect#Of) for a copy.

Some methods must read every value before they yield the first.
They are still lazy: they read when the result is read, not when they are called.

| Reads every value first | Why |
| --- | --- |
| [`SortBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortBy), [`SortFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortFunc), [`Reverse`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Reverse), [`Shuffle`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Shuffle) | The first value is known only at the end |
| [`GroupBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.GroupBy), [`CountBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.CountBy), [`MapToGroups`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MapToGroups) | A group is complete only at the end |
| [`Partition`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Partition) | It reads once for both of its results, and keeps them |
| [`Take(-n)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Take), [`Slice`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Slice) with a negative value, [`Pad(-n, v)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Pad), [`Split`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Split), [`SplitIn`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SplitIn) | They count from the end |
| [`DiffBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffBy), [`IntersectBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.IntersectBy) and the like | They read the other collection whole |

## Key functions and `Identity`

A method cannot add a constraint to `T`, so [`Collection[T]`](https://pkg.go.dev/github.com/mpyw/collect#Collection) has no `Sum()` for numbers only.
Instead, a method that needs numbers, an order or equality takes a key function, and its name ends in `By`.
Pass [`collect.Identity`](https://pkg.go.dev/github.com/mpyw/collect#Identity) when the values are their own keys:

```go
c := collect.Of(3, 1, 2)

c.SumBy(collect.Identity)            // 6
c.SortBy(collect.Identity).All()     // [1 2 3]
c.UniqueBy(collect.Identity).Count() // 3
```

Go infers the type of [`Identity`](https://pkg.go.dev/github.com/mpyw/collect#Identity) from the collection.

| A function that | Name | Examples |
| --- | --- | --- |
| Returns a key to compare, order, group or add up | Ends in `By` | [`SumBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SumBy), [`MaxBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MaxBy), [`UniqueBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.UniqueBy), [`DiffBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffBy), [`SortBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortBy), [`GroupBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.GroupBy), [`ImplodeBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.ImplodeBy), [`WhereInBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.WhereInBy) |
| Returns a bool | No suffix | [`Filter`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Filter), [`Contains`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Contains), [`Every`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Every), [`TakeWhile`](https://pkg.go.dev/github.com/mpyw/collect#Collection.TakeWhile), [`Partition`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Partition) |
| Returns a new value | No suffix | [`Map`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map), [`FlatMap`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap), [`MapWithKeys`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MapWithKeys), [`Reduce`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Reduce) |
| Compares two values | Ends in `Func`, as in the standard library | [`SortFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortFunc), [`DiffFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffFunc), [`IntersectFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.IntersectFunc), [`SortKeysFunc`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.SortKeysFunc) |

## Results: `bool`, `error` and `Must`

| A method returns | When | Examples |
| --- | --- | --- |
| `(T, bool)` | There may be nothing, and that is the only way to fail | [`First`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First), [`Last`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Last), [`Get`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Get), [`Search`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Search), [`MaxBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MaxBy), [`AvgBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.AvgBy) |
| `(T, error)` | It can fail in more than one way, or the reason matters | [`Sole`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sole), [`Ensure`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Ensure), [`Combine`](https://pkg.go.dev/github.com/mpyw/collect#Combine), [`RandomN`](https://pkg.go.dev/github.com/mpyw/collect#Collection.RandomN) |
| `T`, and panics | You know it cannot fail. Its name is `Must` and the name of the method above | [`MustFirst`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustFirst), [`MustLast`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustLast), [`MustGet`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.MustGet), [`MustSole`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustSole), [`MustEnsure`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustEnsure), [`MustCombine`](https://pkg.go.dev/github.com/mpyw/collect#MustCombine) |

Every error wraps one of the package's errors, so [`errors.Is`](https://pkg.go.dev/errors#Is) tells them apart.
A `Must` method panics with the same error.

```go
_, err := collect.Of(1, 2, 3).Sole()
errors.Is(err, collect.ErrMultipleItemsFound) // true
```

> [!IMPORTANT]
> An int argument that cannot be right panics, with no `Must` in the name.
> [`Chunk(0)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk), [`Sliding(0, 1)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sliding), [`Nth(0, 0)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Nth), [`Split(0)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Split) and [`RangeStep(0, 1, 0)`](https://pkg.go.dev/github.com/mpyw/collect#RangeStep) are such calls, as [`slices.Chunk(s, 0)`](https://pkg.go.dev/slices#Chunk) is.

## Keys when you want them

A [`Collection[T]`](https://pkg.go.dev/github.com/mpyw/collect#Collection) has no keys.
[`Keyed()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Keyed) gives each value its index as a key, and returns a [`KeyedCollection[int, T]`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection).
[`KeyBy(f)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.KeyBy) gives each value the key `f` returns.
[`Values()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Values) drops the keys again, and [`Keys()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Keys) keeps only the keys:

```go
odd := collect.Of(10, 11, 12, 13).Keyed().
	Filter(func(i, v int) bool { return v%2 == 1 })

odd.Keys().All()   // [1 3]
odd.Values().All() // [11 13]
```

A [`KeyedCollection[K, V]`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection) wraps an [`iter.Seq2[K, V]`](https://pkg.go.dev/iter#Seq2).
It has every method of [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) but [`Keyed`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Keyed), [`Pad`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Pad), [`Zip`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Zip) and [`CrossJoin`](https://pkg.go.dev/github.com/mpyw/collect#Collection.CrossJoin), and each keeps the keys.
Its functions take the key first, as [`maps.All`](https://pkg.go.dev/maps#All) yields it.
A method that returns a value returns the pair instead: [`First()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First) is `(K, V, bool)`.
It also has the methods of keys: [`Get`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Get), [`Has`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Has), [`Only`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Only), [`Except`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Except), [`FlipBy`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.FlipBy), [`Union`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Union), [`Replace`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Replace), [`DiffKeys`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.DiffKeys), [`SortKeysBy`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.SortKeysBy) and more.

> [!NOTE]
> A key may come more than once, as in an [`iter.Seq2`](https://pkg.go.dev/iter#Seq2).
> [`Count()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Count) and [`Keys()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Keys) see each one, and [`Get`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Get) returns the first value of a key.
> [`All()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.All) returns a `map[K]V`, so the last value of a key wins.

## Nested collections

[`Chunk`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk) returns a [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) of collections.
It is a [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) like any other, so every method applies to it, [`Chunk`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk) included:

```go
quads := collect.Of(1, 2, 3, 4, 5, 6, 7, 8).Chunk(2).Chunk(2)
// [[[1 2] [3 4]] [[5 6] [7 8]]]
```

[`GroupBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.GroupBy), [`ChunkWhile`](https://pkg.go.dev/github.com/mpyw/collect#Collection.ChunkWhile), [`ChunkBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.ChunkBy), [`Sliding`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sliding), [`Split`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Split) and [`SplitIn`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SplitIn) nest the same way.
[`FlatMap(collect.Identity)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap) flattens chunks again.

> [!NOTE]
> In the documentation, [`Chunk`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk) reads [`Chunk[C interface{ Collection[T] }](size int) Collection[C]`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk).
> `C` is always [`Collection[T]`](https://pkg.go.dev/github.com/mpyw/collect#Collection), and Go infers it. Write [`c.Chunk(2)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk).
>
> A method of [`Collection[T]`](https://pkg.go.dev/github.com/mpyw/collect#Collection) cannot return [`Collection[Collection[T]]`](https://pkg.go.dev/github.com/mpyw/collect#Collection) in Go.
> The compiler reports an "instantiation cycle" ([golang/go#80172](https://github.com/golang/go/issues/80172)).
> The type parameter `C` avoids that, and the result is the same type.
> [`Zip`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Zip) and [`CrossJoin`](https://pkg.go.dev/github.com/mpyw/collect#Collection.CrossJoin) return [`Pair`](https://pkg.go.dev/github.com/mpyw/collect#Pair)s the same way.

## JSON

A collection streams itself as JSON as it is read, through [`encoding/json/v2`](https://pkg.go.dev/encoding/json/v2)'s [`MarshalJSONTo`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MarshalJSONTo).
[`json.Marshal`](https://pkg.go.dev/encoding/json#Marshal) of [`encoding/json`](https://pkg.go.dev/encoding/json) uses it too.
A [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) is an array, and a [`KeyedCollection`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection) is an object in its own order:

```go
b, _ := json.Marshal(collect.Of("x", "y").Keyed())
// {"0":"x","1":"y"}
```

## Differences from Laravel

| Laravel | collect | Why |
| --- | --- | --- |
| Keys survive [`filter`](https://laravel.com/framework/docs/collections#method-filter) | Only after [`Keyed()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Keyed) | You choose a [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) or a [`KeyedCollection`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection) |
| `fn($value, $key)` | `func(key, value)` | The order of [`iter.Seq2`](https://pkg.go.dev/iter#Seq2) |
| [`sum()`](https://laravel.com/framework/docs/collections#method-sum), [`sort()`](https://laravel.com/framework/docs/collections#method-sort), [`unique()`](https://laravel.com/framework/docs/collections#method-unique) | [`SumBy(collect.Identity)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SumBy) and the like | See [Key functions and `Identity`](#key-functions-and-identity) |
| [`first()`](https://laravel.com/framework/docs/collections#method-first) returns `null` | [`First()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First) returns `(T, bool)` | Go has no `null` for every type |
| [`firstOrFail()`](https://laravel.com/framework/docs/collections#method-first-or-fail), [`sole()`](https://laravel.com/framework/docs/collections#method-sole) throw | [`MustFirst()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustFirst), [`MustSole()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustSole) panic. [`Sole()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sole) returns an error | See [Results](#results-bool-error-and-must) |
| [`contains($value)`](https://laravel.com/framework/docs/collections#method-contains), [`where('age', '>', 20)`](https://laravel.com/framework/docs/collections#method-where) | [`Contains(func...)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Contains), [`Filter(func...)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Filter) | The compiler checks a function |
| [`pluck('user.name')`](https://laravel.com/framework/docs/collections#method-pluck), [`$c->map->name`](https://laravel.com/framework/docs/collections#higher-order-messages) | [`Map(func...)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map) | No reflection, and no magic properties |
| `*Strict` methods | The methods without `Strict` | Go always compares strictly |
| [`join(', ', ' and ')`](https://laravel.com/framework/docs/collections#method-join) | [`JoinFinalBy(f, ", ", " and ")`](https://pkg.go.dev/github.com/mpyw/collect#Collection.JoinFinalBy) | [`JoinBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.JoinBy) is [`ImplodeBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.ImplodeBy), under the name of [`strings.Join`](https://pkg.go.dev/strings#Join) |
| [`collect($map)`](https://laravel.com/framework/docs/collections#creating-collections) keeps the order | [`FromMap(m)`](https://pkg.go.dev/github.com/mpyw/collect#FromMap) has no order | A Go map has no order. Sort it, or use [`FromSeq2`](https://pkg.go.dev/github.com/mpyw/collect#FromSeq2) |
| [`chunk(0)`](https://laravel.com/framework/docs/collections#method-chunk) is empty | [`Chunk(0)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk) panics | The same as [`slices.Chunk`](https://pkg.go.dev/slices#Chunk) |
| [`zip`](https://laravel.com/framework/docs/collections#method-zip) takes many arrays | [`Zip`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Zip) takes one collection | Go has no variadic type parameters |

Some Laravel methods have another name or another way here:

| Laravel | collect |
| --- | --- |
| [`average`](https://laravel.com/framework/docs/collections#method-average), [`some`](https://laravel.com/framework/docs/collections#method-some) | [`AvgBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.AvgBy), [`Contains`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Contains) |
| [`hasSole()`](https://laravel.com/framework/docs/collections#method-hassole), [`containsOneItem()`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_containsOneItem) | The error of [`Sole()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sole), which returns the value too |
| [`hasMany()`](https://laravel.com/framework/docs/collections#method-hasmany), [`containsManyItems()`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_containsManyItems) | `c.Take(2).Count() == 2` |
| [`merge`](https://laravel.com/framework/docs/collections#method-merge) | [`Concat`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Concat) for a [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection), [`Replace`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Replace) for a [`KeyedCollection`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection) |
| [`mapToDictionary`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_mapToDictionary) | [`MapToGroups`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MapToGroups) |
| [`reduceWithKeys`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_reduceWithKeys) | [`KeyedCollection.Reduce`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Reduce) |
| [`sort($callback)`](https://laravel.com/framework/docs/collections#method-sort), [`sortKeysUsing`](https://laravel.com/framework/docs/collections#method-sortkeysusing), [`diffUsing`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_diffUsing), [`intersectUsing`](https://laravel.com/framework/docs/collections#method-intersectusing) | [`SortFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortFunc), [`SortKeysFunc`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.SortKeysFunc), [`DiffFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffFunc), [`IntersectFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.IntersectFunc) |
| [`value('name')`](https://laravel.com/framework/docs/collections#method-value), [`firstWhere('name', 'x')`](https://laravel.com/framework/docs/collections#method-first-where) | [`Map(f).First()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map), [`FirstWhere(func...)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FirstWhere) |
| [`isEmpty()`](https://laravel.com/framework/docs/collections#method-isempty), [`isNotEmpty()`](https://laravel.com/framework/docs/collections#method-isnotempty) | The `bool` of [`First()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First): `_, ok := c.First()` |
| [`collapse()`](https://laravel.com/framework/docs/collections#method-collapse), [`flatten(1)`](https://laravel.com/framework/docs/collections#method-flatten) | [`FlatMap(collect.FromSlice)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap), [`FlatMap(collect.Identity)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap) |
| [`toArray()`](https://laravel.com/framework/docs/collections#method-toarray), [`getIterator()`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_getIterator), [`toJson()`](https://laravel.com/framework/docs/collections#method-tojson), [`toPrettyJson()`](https://laravel.com/framework/docs/collections#method-to-pretty-json) | [`All()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.All), [`Seq()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Seq), [`json.Marshal(c)`](https://pkg.go.dev/encoding/json/v2#Marshal), [`json.Marshal(c, jsontext.WithIndent("  "))`](https://pkg.go.dev/encoding/json/v2#Marshal) |
| [`make`](https://laravel.com/framework/docs/collections#method-make), [`wrap`](https://laravel.com/framework/docs/collections#method-wrap), [`empty`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_empty), [`fromJson`](https://laravel.com/framework/docs/collections#method-fromjson) | [`Of`](https://pkg.go.dev/github.com/mpyw/collect#Of), [`FromSlice`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice), the zero value, [`json.Unmarshal`](https://pkg.go.dev/encoding/json/v2#Unmarshal) then [`FromSlice`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice) |

These have no Go equivalent, and are not here:

| Laravel | Why |
| --- | --- |
| [`dd`](https://laravel.com/framework/docs/collections#method-dd), [`dump`](https://laravel.com/framework/docs/collections#method-dump) | Use [`fmt`](https://pkg.go.dev/fmt) |
| [`macro`](https://laravel.com/framework/docs/collections#method-macro), [`proxy`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_proxy), [`mapInto`](https://laravel.com/framework/docs/collections#method-mapinto), [`pipeInto`](https://laravel.com/framework/docs/collections#method-pipeinto) | PHP's classes and magic methods |
| [`eachSpread`](https://laravel.com/framework/docs/collections#method-eachspread), [`mapSpread`](https://laravel.com/framework/docs/collections#method-mapspread), [`reduceSpread`](https://laravel.com/framework/docs/collections#method-reduce-spread) | Go has no argument spread |
| [`flatten()`](https://laravel.com/framework/docs/collections#method-flatten) of any depth, [`dot`](https://laravel.com/framework/docs/collections#method-dot), [`undot`](https://laravel.com/framework/docs/collections#method-undot), [`mergeRecursive`](https://laravel.com/framework/docs/collections#method-mergerecursive), [`replaceRecursive`](https://laravel.com/framework/docs/collections#method-replacerecursive) | A Go type has a fixed depth |
| [`whereNull`](https://laravel.com/framework/docs/collections#method-wherenull), [`whereNotNull`](https://laravel.com/framework/docs/collections#method-wherenotnull), [`filter()`](https://laravel.com/framework/docs/collections#method-filter) with no callback | Go has no `null` or falsy value for every type. Pass a function to [`Filter`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Filter) |
| [`diffAssocUsing`](https://laravel.com/framework/docs/collections#method-diffassocusing), [`diffKeysUsing`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_diffKeysUsing), [`intersectAssocUsing`](https://laravel.com/framework/docs/collections#method-intersectassocusing) | Go keys are comparable |
| [`select`](https://laravel.com/framework/docs/collections#method-select), [`escapeWhenCastingToString`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_escapeWhenCastingToString), [`getCachingIterator`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_getCachingIterator) | PHP's arrays and strings |
| [`withHeartbeat`](https://laravel.com/framework/docs/collections#method-with-heartbeat), [`throttle`](https://laravel.com/framework/docs/collections#method-throttle), [`takeUntilTimeout`](https://laravel.com/framework/docs/collections#method-takeUntilTimeout) | They look at the clock only when a value comes, so a slow source stops nothing. A Go program times work with [`context`](https://pkg.go.dev/context), [`time.Ticker`](https://pkg.go.dev/time#Ticker) or [`golang.org/x/time/rate`](https://pkg.go.dev/golang.org/x/time/rate) |
| [`whenEmpty`](https://laravel.com/framework/docs/collections#method-whenempty), [`whenNotEmpty`](https://laravel.com/framework/docs/collections#method-whennotempty), [`unlessEmpty`](https://laravel.com/framework/docs/collections#method-unlessempty), [`unlessNotEmpty`](https://laravel.com/framework/docs/collections#method-unlessnotempty) | They read the first value twice, which breaks a source that can be read once. Use [`First()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First), which returns the value it reads |
| [`tap`](https://laravel.com/framework/docs/collections#method-tap) | To see the whole collection, it must read the source, which may then be empty for the rest of the chain. [`TapEach`](https://pkg.go.dev/github.com/mpyw/collect#Collection.TapEach) sees each value lazily. Use [`Pipe`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Pipe) to call a function in a chain |
| [`reduceInto`](https://laravel.com/framework/docs/collections#method-reduce-into) | PHP copies an array that a callback returns, so Laravel passes the result by reference. A Go map or slice shares its contents, so [`Reduce`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Reduce) does the same |

## License

MIT

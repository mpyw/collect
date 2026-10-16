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

A collection reads its source again on each read, so it is as reusable as its source.
One from [`Of`](https://pkg.go.dev/github.com/mpyw/collect#Of), [`FromSlice`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice) or [`FromMap`](https://pkg.go.dev/github.com/mpyw/collect#FromMap) reads the same values every time, and needs no [`Remember()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Remember):

```go
evens := collect.Of(1, 2, 3, 4).Filter(func(n int) bool { return n%2 == 0 })

evens.All() // [2 4]
evens.All() // [2 4], read again from the start
```

One from [`FromSeq`](https://pkg.go.dev/github.com/mpyw/collect#FromSeq) can be read as many times as its [`iter.Seq`](https://pkg.go.dev/iter#Seq) allows.
A source that can be read only once, such as a database cursor, is used up by the first read.
[`Eager()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Eager) makes a finite collection reusable, and [`Remember()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Remember) any collection:

```go
once := collect.FromSeq(openCursor())
once.All() // [1 2 3]
once.All() // [], the cursor is used up

kept := collect.FromSeq(openCursor()).Remember()
kept.All() // [1 2 3]
kept.All() // [1 2 3], from what Remember kept
```

Go cannot tell these apart by type, as PHP tells a generator function from a [`Generator`](https://www.php.net/manual/en/class.generator.php) object.
It depends on where the [`iter.Seq`](https://pkg.go.dev/iter#Seq) keeps its state:

| An [`iter.Seq`](https://pkg.go.dev/iter#Seq) that | Each call | Like PHP's |
| --- | --- | --- |
| Starts inside the function, as [`slices.Values`](https://pkg.go.dev/slices#Values) reads its slice from the start | Reads again from the start | [Generator function](https://www.php.net/manual/en/language.generators.php) |
| Reads from something opened outside it, such as [`*sql.Rows`](https://pkg.go.dev/database/sql#Rows) | Goes on where the last one stopped, or finds the end | [`Generator`](https://www.php.net/manual/en/class.generator.php) object |

To make a source reusable, open it inside the function, so that each read runs the query again.
To read it only once, call [`Remember()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Remember).

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
> If the source panics, every later read past the kept values panics the same way.

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

A [`KeyedCollection`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection) method passes a key and a value to its function.
Pass [`collect.Value`](https://pkg.go.dev/github.com/mpyw/collect#Value) or [`collect.Key`](https://pkg.go.dev/github.com/mpyw/collect#Key) there, the [`Identity`](https://pkg.go.dev/github.com/mpyw/collect#Identity) of a pair:

```go
prices := collect.FromMap(map[string]int{"desk": 200, "chair": 100})

prices.SumBy(collect.Value)                // 300
prices.SortBy(collect.Key).Keys().All()    // [chair desk]
prices.SortBy(collect.Key).Entries().All() // [{chair 100} {desk 200}]
prices.SortBy(collect.Key).ToMap()         // map[chair:100 desk:200], in no order
```

> [!WARNING]
> [`ToMap()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.ToMap) returns a Go map, so it drops the order that [`SortBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortBy) made.
> A range over the map visits the keys in any order.
> [`fmt`](https://pkg.go.dev/fmt) prints a map sorted by key, which hides this.
> Use [`Entries().All()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Entries) or [`Keys().All()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Keys) to keep the order.

A function that a method also applies to another collection takes only what it compares, so [`Identity`](https://pkg.go.dev/github.com/mpyw/collect#Identity) fits it.
[`DiffBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffBy), [`IntersectBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.IntersectBy), [`DiffAssocBy`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.DiffAssocBy), [`IntersectAssocBy`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.IntersectAssocBy) and [`FlipBy`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.FlipBy) take a value, and [`SortKeysBy`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.SortKeysBy) a key.

| A function that | Name | Examples |
| --- | --- | --- |
| Returns a key to compare, order, group or add up | Ends in `By` | [`SumBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SumBy), [`MaxBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MaxBy), [`UniqueBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.UniqueBy), [`DiffBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffBy), [`SortBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortBy), [`GroupBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.GroupBy), [`ImplodeBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.ImplodeBy), [`WhereInBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.WhereInBy) |
| Returns a bool | No suffix | [`Filter`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Filter), [`Contains`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Contains), [`Every`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Every), [`TakeWhile`](https://pkg.go.dev/github.com/mpyw/collect#Collection.TakeWhile), [`Partition`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Partition) |
| Returns a new value | No suffix | [`Map`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map), [`FlatMap`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap), [`MapWithKeys`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MapWithKeys), [`Reduce`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Reduce) |
| Compares two values | Ends in `Func`, as in the standard library | [`SortFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortFunc), [`DiffFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffFunc), [`IntersectFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.IntersectFunc), [`SortKeysFunc`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.SortKeysFunc) |

## Results: `bool`, `error` and `Must`

<table>
<tr><th>A method returns</th><th>When</th><th>Example</th></tr>
<tr>
<td>

`(T, bool)`

</td>
<td>

There may be nothing, and that is the only way to fail.

[`First`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First), [`Last`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Last), [`Get`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Get), [`Search`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Search), [`MaxBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MaxBy), [`AvgBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.AvgBy)

</td>
<td>

```go
user, ok := users.First()
if !ok {
	// There is no user.
}
```

</td>
</tr>
<tr>
<td>

`(T, error)`

</td>
<td>

It can fail in more than one way, or the reason matters.

[`Sole`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sole), [`Ensure`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Ensure), [`Combine`](https://pkg.go.dev/github.com/mpyw/collect#Combine), [`RandomN`](https://pkg.go.dev/github.com/mpyw/collect#Collection.RandomN)

</td>
<td>

```go
user, err := users.Sole()
if errors.Is(err, collect.ErrMultipleItemsFound) {
	// There is more than one.
}
```

</td>
</tr>
<tr>
<td>

`T`, and panics

</td>
<td>

You know it cannot fail. Its name is `Must` and the name of the method above.

[`MustFirst`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustFirst), [`MustLast`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustLast), [`MustGet`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.MustGet), [`MustSole`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustSole), [`MustEnsure`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustEnsure), [`MustCombine`](https://pkg.go.dev/github.com/mpyw/collect#MustCombine)

</td>
<td>

```go
// Panics with collect.ErrItemNotFound if there is none.
user := users.MustFirst()
```

</td>
</tr>
<tr>
<td>

A panic, with no `Must`

</td>
<td>

An int argument that cannot be right, as [`slices.Chunk(s, 0)`](https://pkg.go.dev/slices#Chunk) panics.

For example [`Chunk`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk), [`Sliding`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sliding), [`Nth`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Nth), [`Split`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Split), [`SplitIn`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SplitIn), [`RangeStep`](https://pkg.go.dev/github.com/mpyw/collect#RangeStep)

</td>
<td>

```go
c.Chunk(0)                 // panics
collect.RangeStep(0, 1, 0) // panics
```

</td>
</tr>
</table>

Every error wraps one of the package's errors, so [`errors.Is`](https://pkg.go.dev/errors#Is) tells them apart.
A `Must` method panics with the same error.

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
[`Entries()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Entries) returns its pairs in order, as a [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) of [`Entry`](https://pkg.go.dev/github.com/mpyw/collect#Entry) values, with every repeated key.

> [!NOTE]
> A key may come more than once, as in an [`iter.Seq2`](https://pkg.go.dev/iter#Seq2).
> [`Count()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Count) and [`Keys()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Keys) see each one, and [`Get`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Get) returns the first value of a key.
> [`ToMap()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.ToMap) returns a `map[K]V`, so it has no order, and the last value of a key wins.
> [`Entries().All()`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Entries) keeps the order and every pair.

## Nested collections

[`Chunk`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk) returns a [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) of collections.
It is a [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) like any other, so every method applies to it, [`Chunk`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk) included:

```go
quads := collect.Of(1, 2, 3, 4, 5, 6, 7, 8).Chunk(2).Chunk(2)

quads.Count()                       // 2
quads.MustFirst().MustFirst().All() // [1 2]
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
> [`Zip`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Zip) and [`CrossJoin`](https://pkg.go.dev/github.com/mpyw/collect#Collection.CrossJoin) return [`Entry`](https://pkg.go.dev/github.com/mpyw/collect#Entry) values the same way.

## JSON

A collection streams itself as JSON as it is read, through [`encoding/json/v2`](https://pkg.go.dev/encoding/json/v2)'s [`MarshalJSONTo`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MarshalJSONTo).
[`json.Marshal`](https://pkg.go.dev/encoding/json#Marshal) of [`encoding/json`](https://pkg.go.dev/encoding/json) uses it too.
A [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) is an array, and a [`KeyedCollection`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection) is an object in its own order:

```go
b, _ := json.Marshal(collect.Of("x", "y").Keyed())
// {"0":"x","1":"y"}
```

> [!NOTE]
> A [`KeyedCollection`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection) is always an object, even where PHP would write an array.
> A key that comes twice is an error in [`encoding/json/v2`](https://pkg.go.dev/encoding/json/v2), and two members in [`encoding/json`](https://pkg.go.dev/encoding/json).

## Go pitfalls

| Pitfall | Do this |
| --- | --- |
| `fmt.Println(c)`, `==` and [`reflect.DeepEqual`](https://pkg.go.dev/reflect#DeepEqual) see a function, not the values | Compare [`c.All()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.All), and range over [`c.Seq()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Seq) |
| A key that is `comparable` but holds a slice in an interface, such as `any`, panics in a map | Return a hashable key from the key function |
| [`FlatMap(collect.Of)`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap) does not compile, as [`Of`](https://pkg.go.dev/github.com/mpyw/collect#Of) is variadic | Pass [`collect.FromSlice`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice) |
| Your own generic type with a method that returns itself over [`Collection[T]`](https://pkg.go.dev/github.com/mpyw/collect#Collection) is an instantiation cycle too | Pin a type parameter, as [`Entries`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Entries) does |

## Differences from Laravel

<table>
<tr><th>Laravel</th><th>collect</th><th>Why</th></tr>
<tr>
<td>

[`filter`](https://laravel.com/framework/docs/collections#method-filter)

```php
collect([10, 11, 12])
    ->filter(fn ($v) => $v > 10)
    ->all();
// [1 => 11, 2 => 12]
```

</td>
<td>

[`Filter`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Filter), [`Keyed`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Keyed), [`Entries`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Entries)

```go
collect.Of(10, 11, 12).
	Filter(func(v int) bool { return v > 10 }).
	All()
// [11 12]

collect.Of(10, 11, 12).Keyed().
	Filter(func(_, v int) bool { return v > 10 }).
	Entries().
	All()
// [{1 11} {2 12}]
```

</td>
<td>

Keys are kept only after [`Keyed()`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Keyed). You choose a [`Collection`](https://pkg.go.dev/github.com/mpyw/collect#Collection) or a [`KeyedCollection`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection)

</td>
</tr>
<tr>
<td>

[`map`](https://laravel.com/framework/docs/collections#method-map)

```php
$c->map(fn ($value, $key) => "$key=$value");
```

</td>
<td>

[`KeyedCollection.Map`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Map)

```go
c.Map(func(key string, value int) string {
	return fmt.Sprintf("%s=%d", key, value)
})
```

</td>
<td>

The key comes first, in the order of [`iter.Seq2`](https://pkg.go.dev/iter#Seq2)

</td>
</tr>
<tr>
<td>

[`sum`](https://laravel.com/framework/docs/collections#method-sum)

```php
$c->sum();
$orders->sum('price');
```

</td>
<td>

[`SumBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SumBy)

```go
c.SumBy(collect.Identity)
orders.SumBy(func(o Order) int { return o.Price })
```

</td>
<td>

See [Key functions and [`Identity`](https://pkg.go.dev/github.com/mpyw/collect#Identity)](#key-functions-and-identity)

</td>
</tr>
<tr>
<td>

[`first`](https://laravel.com/framework/docs/collections#method-first)

```php
$user = $users->first(); // null if none
```

</td>
<td>

[`First`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First)

```go
user, ok := users.First()
```

</td>
<td>

Go has no `null` for every type

</td>
</tr>
<tr>
<td>

[`firstOrFail`](https://laravel.com/framework/docs/collections#method-first-or-fail), [`sole`](https://laravel.com/framework/docs/collections#method-sole)

```php
$users->firstOrFail(); // throws
$users->sole();        // throws
```

</td>
<td>

[`MustFirst`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustFirst), [`MustSole`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MustSole), [`Sole`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sole)

```go
users.MustFirst() // panics
users.MustSole()  // panics
users.Sole()      // returns an error
```

</td>
<td>

See [Results](#results-bool-error-and-must)

</td>
</tr>
<tr>
<td>

[`contains`](https://laravel.com/framework/docs/collections#method-contains), [`where`](https://laravel.com/framework/docs/collections#method-where)

```php
$c->contains(3);
$users->where('age', '>', 20);
```

</td>
<td>

[`Contains`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Contains), [`Filter`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Filter)

```go
c.Contains(func(v int) bool { return v == 3 })
users.Filter(func(u User) bool { return u.Age > 20 })
```

</td>
<td>

The compiler checks a function

</td>
</tr>
<tr>
<td>

[`pluck`](https://laravel.com/framework/docs/collections#method-pluck), [higher order messages](https://laravel.com/framework/docs/collections#higher-order-messages)

```php
$users->pluck('name');
$users->map->name;
```

</td>
<td>

[`Map`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map)

```go
users.Map(func(u User) string { return u.Name })
```

</td>
<td>

No reflection, and no magic properties

</td>
</tr>
<tr>
<td>

`*Strict` methods

</td>
<td>

The methods without `Strict`

</td>
<td>

Go always compares strictly

</td>
</tr>
<tr>
<td>

[`join`](https://laravel.com/framework/docs/collections#method-join)

```php
collect(['a', 'b', 'c'])->join(', ', ' and ');
// 'a, b and c'
```

</td>
<td>

[`JoinFinalBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.JoinFinalBy)

```go
collect.Of("a", "b", "c").
	JoinFinalBy(collect.Identity, ", ", " and ")
// "a, b and c"
```

</td>
<td>

[`JoinBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.JoinBy) is [`ImplodeBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.ImplodeBy), under the name of [`strings.Join`](https://pkg.go.dev/strings#Join)

</td>
</tr>
<tr>
<td>

[`collect($map)`](https://laravel.com/framework/docs/collections#creating-collections), [`keys`](https://laravel.com/framework/docs/collections#method-keys)

```php
collect(['b' => 2, 'a' => 1])->keys();
// ['b', 'a']
```

</td>
<td>

[`FromMap`](https://pkg.go.dev/github.com/mpyw/collect#FromMap), [`Keys`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Keys)

```go
collect.FromMap(map[string]int{"b": 2, "a": 1}).Keys()
// either order
```

</td>
<td>

A Go map has no order. Sort it, or use [`FromSeq2`](https://pkg.go.dev/github.com/mpyw/collect#FromSeq2)

</td>
</tr>
<tr>
<td>

[`chunk`](https://laravel.com/framework/docs/collections#method-chunk)

```php
$c->chunk(0); // empty
```

</td>
<td>

[`Chunk`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Chunk)

```go
c.Chunk(0) // panics
```

</td>
<td>

The same as [`slices.Chunk`](https://pkg.go.dev/slices#Chunk)

</td>
</tr>
<tr>
<td>

[`zip`](https://laravel.com/framework/docs/collections#method-zip)

```php
$a->zip($b, $c);
```

</td>
<td>

[`Zip`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Zip)

```go
a.Zip(b) // one at a time
```

</td>
<td>

Go has no variadic type parameters

</td>
</tr>
<tr>
<td>

[`skip`](https://laravel.com/framework/docs/collections#method-skip)

```php
$lazy->skip(-1); // empty
```

</td>
<td>

[`Skip`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Skip)

```go
c.Skip(-1) // every value
```

</td>
<td>

A negative count is 0. Laravel's lazy and eager skip disagree here

</td>
</tr>
<tr>
<td>

[`percentage`](https://laravel.com/framework/docs/collections#method-percentage)

```php
$c->percentage($f, 2);
```

</td>
<td>

[`Percentage`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Percentage)

```go
c.Percentage(f, 2)
```

</td>
<td>

PHP's `round` works on the decimal form, Go's [`math.Round`](https://pkg.go.dev/math#Round) on the binary one. A tie can round the other way

</td>
</tr>
<tr>
<td>

[`nth`](https://laravel.com/framework/docs/collections#method-nth), [`multiply`](https://laravel.com/framework/docs/collections#method-multiply), [`shuffle`](https://laravel.com/framework/docs/collections#method-shuffle), [`random`](https://laravel.com/framework/docs/collections#method-random), [`split`](https://laravel.com/framework/docs/collections#method-split), [`flatMap`](https://laravel.com/framework/docs/collections#method-flatmap)

```php
$keyed->nth(2); // keys 0, 1, 2
```

</td>
<td>

[`Nth`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Nth), [`Multiply`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Multiply), [`Shuffle`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Shuffle), [`RandomN`](https://pkg.go.dev/github.com/mpyw/collect#Collection.RandomN), [`Split`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Split), [`FlatMap`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap)

```go
keyed.Nth(2, 0) // keys 0, 2, 4
```

</td>
<td>

A [`KeyedCollection`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection) keeps every key, where Laravel numbers them again

</td>
</tr>
<tr>
<td>

[`diffUsing`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_diffUsing), [`intersectUsing`](https://laravel.com/framework/docs/collections#method-intersectusing)

```php
$c->diffUsing($other, fn ($a, $b) => strcasecmp($a, $b));
```

</td>
<td>

[`DiffFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffFunc), [`IntersectFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.IntersectFunc)

```go
c.DiffFunc(other, strings.EqualFold)
```

</td>
<td>

The function tells whether two values are equal, not how they compare

</td>
</tr>
</table>

Some Laravel methods have another name or another way here:

<table>
<tr><th>Laravel</th><th>collect</th></tr>
<tr>
<td>

[`average`](https://laravel.com/framework/docs/collections#method-average), [`some`](https://laravel.com/framework/docs/collections#method-some)

```php
$c->average();
$c->some(fn ($v) => $v > 1);
```

</td>
<td>

[`AvgBy`](https://pkg.go.dev/github.com/mpyw/collect#Collection.AvgBy), [`Contains`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Contains)

```go
c.AvgBy(collect.Identity)
c.Contains(func(v int) bool { return v > 1 })
```

</td>
</tr>
<tr>
<td>

[`hasSole`](https://laravel.com/framework/docs/collections#method-hassole), [`containsOneItem`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_containsOneItem)

```php
$c->hasSole();
$c->containsOneItem();
```

</td>
<td>

[`Sole`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Sole)

```go
_, err := c.Sole() // returns the value too
```

</td>
</tr>
<tr>
<td>

[`hasMany`](https://laravel.com/framework/docs/collections#method-hasmany), [`containsManyItems`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_containsManyItems)

```php
$c->hasMany();
$c->containsManyItems();
```

</td>
<td>

[`Take`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Take), [`Count`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Count)

```go
c.Take(2).Count() == 2
```

</td>
</tr>
<tr>
<td>

[`merge`](https://laravel.com/framework/docs/collections#method-merge)

```php
$list->merge($other);
$keyed->merge($other);
```

</td>
<td>

[`Concat`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Concat), [`KeyedCollection.Replace`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Replace)

```go
list.Concat(other)
keyed.Replace(other)
```

</td>
</tr>
<tr>
<td>

[`mapToDictionary`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_mapToDictionary)

</td>
<td>

[`MapToGroups`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MapToGroups)

</td>
</tr>
<tr>
<td>

[`reduceWithKeys`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_reduceWithKeys)

```php
$c->reduceWithKeys(
    fn ($carry, $value, $key) => $carry + $value,
    0,
);
```

</td>
<td>

[`KeyedCollection.Reduce`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.Reduce)

```go
keyed.Reduce(func(carry int, key string, value int) int {
	return carry + value
}, 0)
```

</td>
</tr>
<tr>
<td>

[`sort($callback)`](https://laravel.com/framework/docs/collections#method-sort), [`sortKeysUsing`](https://laravel.com/framework/docs/collections#method-sortkeysusing), [`diffUsing`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_diffUsing), [`intersectUsing`](https://laravel.com/framework/docs/collections#method-intersectusing)

</td>
<td>

[`SortFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.SortFunc), [`SortKeysFunc`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.SortKeysFunc), [`DiffFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.DiffFunc), [`IntersectFunc`](https://pkg.go.dev/github.com/mpyw/collect#Collection.IntersectFunc)

</td>
</tr>
<tr>
<td>

[`value`](https://laravel.com/framework/docs/collections#method-value), [`firstWhere`](https://laravel.com/framework/docs/collections#method-first-where)

```php
$users->value('name');
$users->firstWhere('name', 'x');
```

</td>
<td>

[`Map`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map), [`First`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First), [`FirstWhere`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FirstWhere)

```go
users.Map(func(u User) string { return u.Name }).First()
users.FirstWhere(func(u User) bool { return u.Name == "x" })
```

</td>
</tr>
<tr>
<td>

[`isEmpty`](https://laravel.com/framework/docs/collections#method-isempty), [`isNotEmpty`](https://laravel.com/framework/docs/collections#method-isnotempty)

```php
$c->isEmpty();
$c->isNotEmpty();
```

</td>
<td>

[`First`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First)

```go
_, ok := c.First()
```

</td>
</tr>
<tr>
<td>

[`collapse`](https://laravel.com/framework/docs/collections#method-collapse), [`flatten`](https://laravel.com/framework/docs/collections#method-flatten)

```php
collect([[1, 2], [3]])->collapse();
$chunks->flatten(1);
```

</td>
<td>

[`FlatMap`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap), [`FromSlice`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice), [`Identity`](https://pkg.go.dev/github.com/mpyw/collect#Identity)

```go
collect.Of([]int{1, 2}, []int{3}).FlatMap(collect.FromSlice)
chunks.FlatMap(collect.Identity)
```

</td>
</tr>
<tr>
<td>

[`toArray`](https://laravel.com/framework/docs/collections#method-toarray), [`getIterator`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_getIterator), [`toJson`](https://laravel.com/framework/docs/collections#method-tojson), [`toPrettyJson`](https://laravel.com/framework/docs/collections#method-to-pretty-json)

```php
$c->toArray();
$c->getIterator();
$c->toJson();
$c->toPrettyJson();
```

</td>
<td>

[`All`](https://pkg.go.dev/github.com/mpyw/collect#Collection.All), [`Seq`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Seq), [`MarshalJSONTo`](https://pkg.go.dev/github.com/mpyw/collect#Collection.MarshalJSONTo)

```go
c.All()
c.Seq()
json.Marshal(c) // encoding/json/v2
json.Marshal(c, jsontext.WithIndent("  "))
```

</td>
</tr>
<tr>
<td>

[`make`](https://laravel.com/framework/docs/collections#method-make), [`wrap`](https://laravel.com/framework/docs/collections#method-wrap), [`empty`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_empty), [`fromJson`](https://laravel.com/framework/docs/collections#method-fromjson)

```php
collect([1, 2]);
Collection::wrap($array);
Collection::empty();
Collection::fromJson('[1, 2]');
```

</td>
<td>

[`Of`](https://pkg.go.dev/github.com/mpyw/collect#Of), [`FromSlice`](https://pkg.go.dev/github.com/mpyw/collect#FromSlice)

```go
collect.Of(1, 2)
collect.FromSlice(slice)
var c collect.Collection[int] // the zero value is empty

var values []int
json.Unmarshal([]byte("[1, 2]"), &values)
collect.FromSlice(values)
```

</td>
</tr>
<tr>
<td>

[`intersectByKeys`](https://laravel.com/framework/docs/collections#method-intersectbykeys)

</td>
<td>

[`IntersectKeys`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.IntersectKeys), named as [`DiffKeys`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.DiffKeys) is

</td>
</tr>
</table>

These have no Go equivalent, and are not here:

<table>
<tr><th>Laravel</th><th>Why</th><th>In Go</th></tr>
<tr>
<td>

[`dd`](https://laravel.com/framework/docs/collections#method-dd), [`dump`](https://laravel.com/framework/docs/collections#method-dump)

</td>
<td>

PHP debugging helpers

</td>
<td>

[`fmt`](https://pkg.go.dev/fmt)

```go
fmt.Println(c.All())
```

</td>
</tr>
<tr>
<td>

[`macro`](https://laravel.com/framework/docs/collections#method-macro), [`proxy`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_proxy), [`mapInto`](https://laravel.com/framework/docs/collections#method-mapinto), [`pipeInto`](https://laravel.com/framework/docs/collections#method-pipeinto)

</td>
<td>

PHP's classes and magic methods

</td>
<td>

[`Pipe`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Pipe), [`Map`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map)

```go
c.Pipe(f)
c.Map(NewThing)
```

</td>
</tr>
<tr>
<td>

[`eachSpread`](https://laravel.com/framework/docs/collections#method-eachspread), [`mapSpread`](https://laravel.com/framework/docs/collections#method-mapspread), [`reduceSpread`](https://laravel.com/framework/docs/collections#method-reduce-spread)

</td>
<td>

Go has no argument spread

</td>
<td>

Take a struct or an [`Entry`](https://pkg.go.dev/github.com/mpyw/collect#Entry)

</td>
</tr>
<tr>
<td>

[`flatten()`](https://laravel.com/framework/docs/collections#method-flatten) of any depth, [`dot`](https://laravel.com/framework/docs/collections#method-dot), [`undot`](https://laravel.com/framework/docs/collections#method-undot), [`mergeRecursive`](https://laravel.com/framework/docs/collections#method-mergerecursive), [`replaceRecursive`](https://laravel.com/framework/docs/collections#method-replacerecursive)

</td>
<td>

A Go type has a fixed depth

</td>
<td>

[`FlatMap`](https://pkg.go.dev/github.com/mpyw/collect#Collection.FlatMap) one level at a time

</td>
</tr>
<tr>
<td>

[`whereNull`](https://laravel.com/framework/docs/collections#method-wherenull), [`whereNotNull`](https://laravel.com/framework/docs/collections#method-wherenotnull), [`filter()`](https://laravel.com/framework/docs/collections#method-filter) with no callback

</td>
<td>

Go has no `null` or falsy value for every type

</td>
<td>

[`Filter`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Filter)

```go
users.Filter(func(u *User) bool { return u != nil })
```

</td>
</tr>
<tr>
<td>

[`diffAssocUsing`](https://laravel.com/framework/docs/collections#method-diffassocusing), [`diffKeysUsing`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_diffKeysUsing), [`intersectAssocUsing`](https://laravel.com/framework/docs/collections#method-intersectassocusing)

</td>
<td>

Go keys are comparable

</td>
<td>

[`DiffAssocBy`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.DiffAssocBy), [`DiffKeys`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.DiffKeys), [`IntersectAssocBy`](https://pkg.go.dev/github.com/mpyw/collect#KeyedCollection.IntersectAssocBy)

</td>
</tr>
<tr>
<td>

[`select`](https://laravel.com/framework/docs/collections#method-select), [`escapeWhenCastingToString`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_escapeWhenCastingToString), [`getCachingIterator`](https://api.laravel.com/docs/master/Illuminate/Support/LazyCollection.html#method_getCachingIterator)

</td>
<td>

PHP's arrays and strings

</td>
<td>

[`Map`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Map)

</td>
</tr>
<tr>
<td>

[`withHeartbeat`](https://laravel.com/framework/docs/collections#method-with-heartbeat), [`throttle`](https://laravel.com/framework/docs/collections#method-throttle), [`takeUntilTimeout`](https://laravel.com/framework/docs/collections#method-takeUntilTimeout)

</td>
<td>

They look at the clock only when a value comes, so a slow source stops nothing

</td>
<td>

[`TakeWhile`](https://pkg.go.dev/github.com/mpyw/collect#Collection.TakeWhile), [`context`](https://pkg.go.dev/context)

```go
c.TakeWhile(func(int) bool { return ctx.Err() == nil })
```

Or [`time.Ticker`](https://pkg.go.dev/time#Ticker), [`golang.org/x/time/rate`](https://pkg.go.dev/golang.org/x/time/rate)

</td>
</tr>
<tr>
<td>

[`whenEmpty`](https://laravel.com/framework/docs/collections#method-whenempty), [`whenNotEmpty`](https://laravel.com/framework/docs/collections#method-whennotempty), [`unlessEmpty`](https://laravel.com/framework/docs/collections#method-unlessempty), [`unlessNotEmpty`](https://laravel.com/framework/docs/collections#method-unlessnotempty)

</td>
<td>

They read the first value twice, which breaks a source that can be read once

</td>
<td>

[`First`](https://pkg.go.dev/github.com/mpyw/collect#Collection.First)

```go
first, ok := c.First()
```

</td>
</tr>
<tr>
<td>

[`tap`](https://laravel.com/framework/docs/collections#method-tap)

</td>
<td>

To see the whole collection, it must read the source, which may then be empty for the rest of the chain

</td>
<td>

[`TapEach`](https://pkg.go.dev/github.com/mpyw/collect#Collection.TapEach)

```go
c.TapEach(func(v int) { fmt.Println(v) })
```

</td>
</tr>
<tr>
<td>

[`reduceInto`](https://laravel.com/framework/docs/collections#method-reduce-into)

</td>
<td>

PHP copies an array that a callback returns. A Go map or slice shares its contents

</td>
<td>

[`Reduce`](https://pkg.go.dev/github.com/mpyw/collect#Collection.Reduce)

```go
c.Reduce(func(m map[int]int, v int) map[int]int {
	m[v]++
	return m
}, map[int]int{})
```

</td>
</tr>
</table>

## License

MIT

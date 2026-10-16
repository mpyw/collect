// Package collect chains lazy operations on a sequence of values, in the
// manner of Laravel's LazyCollection.
//
// A [Collection] wraps an [iter.Seq]. Its methods build a new sequence and
// run nothing: the values flow only when a terminal method, such as
// [Collection.All], or a range over [Collection.Seq] reads them.
//
//declscope:core
//declscope:shared // every method file builds these types
package collect

import (
	"errors"
	"fmt"
	"iter"
	"maps"
	"reflect"
	"slices"
)

// Collection is a lazy sequence of values of type T.
//
// A method whose result nests T, such as [Collection.Chunk], declares a type
// parameter C that only T's collection satisfies, and returns Collection[C].
// Go infers C, so a caller never writes it. Returning
// Collection[Collection[T]] itself would not compile: the compiler would
// instantiate the methods of Collection[Collection[T]], and so of
// Collection[Collection[Collection[T]]], without end ("instantiation
// cycle", golang/go#80172). C is instantiated only by a call, so only as
// deep as calls go.
//
// The zero value is an empty collection. A collection is as reusable as the
// sequence it reads: one from [Of] or [FromSlice] can be read any number of
// times, one from [FromSeq] as many times as its sequence allows.
type Collection[T any] struct {
	seq iter.Seq[T]
}

// KeyedCollection is a lazy sequence of values of type V, each with a key
// of type K. [Collection.Keyed] and [Collection.KeyBy] return it, and
// [KeyedCollection.Values] turns it back into a [Collection].
//
// [FromSeq2] and [FromMap] make one from the standard library's types.
//
// It is a sequence of pairs, not a map: a key may come more than once.
// Only [KeyedCollection.ToMap] collapses them, as a map does.
type KeyedCollection[K comparable, V any] struct {
	seq iter.Seq2[K, V]
}

// Of returns a collection of values. It copies values, so changing the
// slice later does not change the collection. [FromSlice] reads a slice
// without a copy.
func Of[T any](values ...T) Collection[T] {
	return FromSlice(slices.Clone(values))
}

// FromSlice returns a collection that reads s, with no copy. Each read sees
// s as it is then, as [slices.Values] does. [Of] copies instead.
func FromSlice[T any](s []T) Collection[T] {
	return Collection[T]{seq: slices.Values(s)}
}

// FromSeq returns a collection that reads seq. A nil seq is empty.
func FromSeq[T any](seq iter.Seq[T]) Collection[T] {
	return Collection[T]{seq: seq}
}

// FromSeq2 returns a keyed collection that reads seq. A nil seq is empty.
func FromSeq2[K comparable, V any](seq iter.Seq2[K, V]) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: seq}
}

// FromMap returns a keyed collection that reads m, with no copy. As with
// range over a map, the order of the pairs is not specified, and differs
// from one read to the next.
func FromMap[K comparable, V any](m map[K]V) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: maps.All(m)}
}

// Identity returns v. Pass it where a method takes a key function and the
// values are their own keys: c.SumBy(collect.Identity) sums a Collection[int],
// as Laravel's sum() does. Go infers its type from the collection.
func Identity[T any](v T) T {
	return v
}

// Value returns v. It is [Identity] for a [KeyedCollection], whose methods
// take a key and a value: k.SumBy(collect.Value) sums the values.
func Value[K, V any](_ K, v V) V {
	return v
}

// Key returns k. It is [Identity] for the keys of a [KeyedCollection]:
// k.SortBy(collect.Key) sorts the pairs by key.
func Key[K, V any](k K, _ V) K {
	return k
}

// Number is the constraint of the values that [Collection.SumBy],
// [Collection.AvgBy], [Collection.MedianBy], [Range] and [RangeStep]
// compute with.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Entry holds a key and its value, as an entry of an ECMAScript Map does.
// [KeyedCollection.Entries] returns the entries of a keyed collection, and
// [Collection.Entries] those of a collection with its indexes as keys.
type Entry[K, V any] struct {
	Key   K
	Value V
}

// Tuple2 holds a value of each of two collections. [Collection.Zip2] and
// [Collection.CrossJoin] return them. It is not an [Entry]: neither value
// is the key of the other.
type Tuple2[A, B any] struct {
	A A
	B B
}

// Tuple3 holds a value of each of three collections. [Collection.Zip3]
// returns them.
type Tuple3[A, B, C any] struct {
	A A
	B B
	C C
}

// Tuple4 holds a value of each of four collections. [Collection.Zip4]
// returns them.
type Tuple4[A, B, C, D any] struct {
	A A
	B B
	C C
	D D
}

// Tuple5 holds a value of each of five collections. [Collection.Zip5]
// returns them.
type Tuple5[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

var (
	// ErrItemNotFound is the error of [Collection.Sole] when there is no
	// value, and the panic of [Collection.MustFirst], [Collection.MustLast],
	// [KeyedCollection.MustGet] and the like. It is Laravel's
	// ItemNotFoundException.
	ErrItemNotFound = errors.New("collect: item not found")

	// ErrMultipleItemsFound is the error of [Collection.Sole], and the panic
	// of [Collection.MustSole], when there is more than one value. It is
	// Laravel's MultipleItemsFoundException.
	ErrMultipleItemsFound = errors.New("collect: multiple items found")

	// ErrNotEnoughItems is the error of [Collection.RandomN] when it is
	// asked for more values than there are.
	ErrNotEnoughItems = errors.New("collect: not enough items")

	// ErrInvalidArgument is the panic of a method given an int argument
	// that cannot be right, such as Chunk(0), as slices.Chunk panics.
	ErrInvalidArgument = errors.New("collect: invalid argument")

	// ErrUnsupportedKey is the error of [KeyedCollection.MarshalJSONTo] for
	// a key that encodes as neither a JSON string, number nor boolean.
	ErrUnsupportedKey = errors.New("collect: unsupported key")

	// ErrUnexpectedValue is the error of [Collection.Ensure], and the panic
	// of [Collection.MustEnsure], for a value of another type. It is
	// Laravel's UnexpectedValueException.
	ErrUnexpectedValue = errors.New("collect: unexpected value")

	// ErrLengthMismatch is the error of [Combine], and the panic of
	// [MustCombine], when there are not as many values as keys. It is the
	// ValueError of Laravel's combine.
	ErrLengthMismatch = errors.New("collect: lengths do not match")
)

// asType returns v as U, or an error that wraps ErrUnexpectedValue and
// names both types.
func asType[T, U any](v T) (U, error) {
	u, ok := any(v).(U)
	if !ok {
		return u, fmt.Errorf("%w: a value of type %T is not %v", ErrUnexpectedValue, v, reflect.TypeFor[U]())
	}
	return u, nil
}

// invalidArgument panics with an error that wraps ErrInvalidArgument.
func invalidArgument(format string, args ...any) {
	panic(fmt.Errorf("%w: "+format, append([]any{ErrInvalidArgument}, args...)...))
}

// pairsOf reads every pair of c into a slice of keys and one of values.
func pairsOf[K comparable, V any](c KeyedCollection[K, V]) ([]K, []V) {
	var (
		ks []K
		vs []V
	)
	for k, v := range c.Seq() {
		ks, vs = append(ks, k), append(vs, v)
	}
	return ks, vs
}

// keyedOf returns a keyed collection that reads ks and vs, with no copy.
func keyedOf[K comparable, V any](ks []K, vs []V) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for i, k := range ks {
			if !yield(k, vs[i]) {
				return
			}
		}
	}}
}

// groupsOf groups what each adds by key, in the order the keys first come.
// It finds a group by its position, not by looking its key up again, so a
// key that never equals itself, a NaN, is a group of its own each time
// rather than lost.
func groupsOf[K comparable, E any](each func(add func(K, E))) ([]K, [][]E) {
	index := map[K]int{}
	var (
		keys   []K
		groups [][]E
	)
	each(func(k K, e E) {
		i, ok := index[k]
		if !ok {
			i = len(keys)
			index[k] = i
			keys, groups = append(keys, k), append(groups, nil)
		}
		groups[i] = append(groups[i], e)
	})
	return keys, groups
}

package collect_test

import (
	"errors"
	"iter"
	"slices"
	"testing"

	"github.com/mpyw/collect"
)

//declscope:shared // the examples and the method tests
type user struct {
	Name string
	Age  int
}

// counted returns 1, 2, 3, ... without end, and a pointer to how many
// values it has yielded, so that a test can see what a chain read.
//
//declscope:shared // every method test reads an endless source
func counted() (iter.Seq[int], *int) {
	n := 0
	return func(yield func(int) bool) {
		for {
			n++
			if !yield(n) {
				return
			}
		}
	}, &n
}

// endless returns a collection of counted values, for a test that does not
// look at the count.
//
//declscope:shared // every method test reads an endless source
func endless() collect.Collection[int] {
	seq, _ := counted()
	return collect.FromSeq(seq)
}

//declscope:shared // every method test compares slices
func assertEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// assertPairs checks the keys and the values of c, in order.
//
//declscope:shared // the tests of the methods of KeyedCollection
func assertPairs[K comparable, V comparable](t *testing.T, c collect.KeyedCollection[K, V], keys []K, values []V) {
	t.Helper()
	var ks []K
	var vs []V
	for k, v := range c.Seq() {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	assertEqual(t, ks, keys)
	assertEqual(t, vs, values)
}

// assertStopsEarly ranges over seq and breaks after 3 values. A lazy method
// that calls yield again after it returned false makes the range panic.
//
//declscope:shared // every lazy method's test
func assertStopsEarly[T any](t *testing.T, seq iter.Seq[T]) {
	t.Helper()
	n := 0
	for range seq {
		if n++; n == 3 {
			return
		}
	}
	t.Errorf("the sequence ended after %d values, want at least 3", n)
}

// assertStopsEarly2 is assertStopsEarly for an iter.Seq2.
//
//declscope:shared // every lazy method's test
func assertStopsEarly2[K, V any](t *testing.T, seq iter.Seq2[K, V]) {
	t.Helper()
	n := 0
	for range seq {
		if n++; n == 3 {
			return
		}
	}
	t.Errorf("the sequence ended after %d pairs, want at least 3", n)
}

// assertChunks checks a collection of collections against want.
//
//declscope:shared // the tests of the methods that return chunks
func assertChunks[T comparable](t *testing.T, c collect.Collection[collect.Collection[T]], want [][]T) {
	t.Helper()
	got := c.Map(collect.Collection[T].All).All()
	if !slices.EqualFunc(got, want, slices.Equal) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// assertPanics checks that f panics.
//
//declscope:shared // the tests of the methods that panic on a bad argument
func assertPanics(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Error("did not panic")
		}
	}()
	f()
}

// letters returns a collection of each letter of s, as a string.
//
//declscope:shared // the tests that use Laravel's examples of letters
func letters(s string) collect.Collection[string] {
	return collect.FromSeq(func(yield func(string) bool) {
		for _, r := range s {
			if !yield(string(r)) {
				return
			}
		}
	})
}

// keyedOf returns a keyed collection of the keys and values, in order.
//
//declscope:shared // the tests of the methods of KeyedCollection
func keyedOf[K comparable, V any](keys []K, values []V) collect.KeyedCollection[K, V] {
	return collect.MustCombine(collect.FromSlice(keys), collect.FromSlice(values))
}

// once returns a sequence that yields values on its first read only, as a
// network stream or a database cursor would.
//
//declscope:shared // the tests of Eager and Remember
func once[T any](values ...T) collect.Collection[T] {
	read := false
	return collect.FromSeq(func(yield func(T) bool) {
		if read {
			return
		}
		read = true
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	})
}

//declscope:shared // the examples of the Where methods
type priced struct {
	Product string
	Price   int
}

// assertPanicsWith checks that f panics with an error that is target.
//
//declscope:shared // the tests of the Must methods
func assertPanicsWith(t *testing.T, target error, f func()) {
	t.Helper()
	defer func() {
		err, ok := recover().(error)
		if !ok || !errors.Is(err, target) {
			t.Errorf("panicked with %v, want %v", err, target)
		}
	}()
	f()
}

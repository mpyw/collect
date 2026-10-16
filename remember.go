package collect

import (
	"iter"
	"runtime"
)

// Remember returns a collection that reads the source once, and keeps
// each value it reads. A later read gets the kept values first, then reads
// on from where the source stopped. So a source that can be read only once
// can be read again.
//
// It is not safe for concurrent reads. The source stays open between reads
// until it ends, or until the collection is garbage collected; a source that
// refers to the collection keeps it alive, and stays open. If the source
// panics, every later read past the kept values panics the same way. The
// source must not read the collection past the kept values: that panics
// too, as iter.Pull does.
func (c Collection[T]) Remember() Collection[T] {
	r := &rememberState[T]{source: c.Seq()}
	return Collection[T]{seq: func(yield func(T) bool) {
		for i := 0; ; i++ {
			v, ok := r.at(i)
			if !ok || !yield(v) {
				return
			}
		}
	}}
}

// Remember returns a collection that reads the source once, and keeps each
// pair it reads, as [Collection.Remember] does.
func (c KeyedCollection[K, V]) Remember() KeyedCollection[K, V] {
	// Not a Collection[Entry[K, V]]: a method of KeyedCollection that names
	// one, rather than pin it, is an instantiation cycle.
	r := &rememberState[Entry[K, V]]{source: func(yield func(Entry[K, V]) bool) {
		for k, v := range c.Seq() {
			if !yield(Entry[K, V]{k, v}) {
				return
			}
		}
	}}
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for i := 0; ; i++ {
			p, ok := r.at(i)
			if !ok || !yield(p.Key, p.Value) {
				return
			}
		}
	}}
}

type rememberState[T any] struct {
	source   iter.Seq[T]
	next     func() (T, bool)
	kept     []T
	done     bool
	panicked any
}

// at returns the ith value, from the kept ones, or read from the source.
func (r *rememberState[T]) at(i int) (T, bool) {
	if i < len(r.kept) {
		return r.kept[i], true
	}
	var zero T
	if r.panicked != nil {
		panic(r.panicked)
	}
	if r.done {
		return zero, false
	}
	if r.next == nil {
		next, stop := iter.Pull(r.source)
		r.next = next
		// A source not read to its end is stopped when no reader is left.
		runtime.AddCleanup(r, func(stop func()) { stop() }, stop)
	}
	v, ok := r.pull()
	if !ok {
		r.done = true
		return zero, false
	}
	r.kept = append(r.kept, v)
	return v, true
}

// pull reads the next value of the source. If the source panics, iter.Pull
// reports its end to every later call, so pull keeps the panic and the
// source is never read again: a later read past the kept values panics the
// same way, rather than end early.
func (r *rememberState[T]) pull() (T, bool) {
	defer func() {
		if p := recover(); p != nil {
			r.panicked = p
			panic(p)
		}
	}()
	return r.next()
}

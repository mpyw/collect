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
// until it ends, or until the collection is garbage collected.
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
	// Not a Collection[Pair[K, V]]: a method of KeyedCollection that names
	// one is an instantiation cycle, through Collection.KeyBy.
	r := &rememberState[Pair[K, V]]{source: func(yield func(Pair[K, V]) bool) {
		for k, v := range c.Seq() {
			if !yield(Pair[K, V]{k, v}) {
				return
			}
		}
	}}
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for i := 0; ; i++ {
			p, ok := r.at(i)
			if !ok || !yield(p.First, p.Second) {
				return
			}
		}
	}}
}

type rememberState[T any] struct {
	source iter.Seq[T]
	next   func() (T, bool)
	kept   []T
	done   bool
}

// at returns the ith value, from the kept ones, or read from the source.
func (r *rememberState[T]) at(i int) (T, bool) {
	if i < len(r.kept) {
		return r.kept[i], true
	}
	var zero T
	if r.done {
		return zero, false
	}
	if r.next == nil {
		next, stop := iter.Pull(r.source)
		r.next = next
		// A source not read to its end is stopped when no reader is left.
		runtime.AddCleanup(r, func(stop func()) { stop() }, stop)
	}
	v, ok := r.next()
	if !ok {
		r.done = true
		return zero, false
	}
	r.kept = append(r.kept, v)
	return v, true
}

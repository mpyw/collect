package collect

import "sync"

// Partition returns the values for which match returns true, and those for
// which it returns false. It reads the collection once, when either result
// is first read, and keeps both, as Laravel's partition reads once. Later
// reads of either return the kept values.
//
// If match or the source panics, nothing is kept, and the next read starts
// again. match must not read either result: it would wait for itself.
func (c Collection[T]) Partition(match func(T) bool) (Collection[T], Collection[T]) {
	s := &partitionState[T]{}
	split := func() {
		s.split(func(add func(T, bool)) {
			for v := range c.Seq() {
				add(v, match(v))
			}
		})
	}
	half := func(values *[]T) Collection[T] {
		return Collection[T]{seq: func(yield func(T) bool) {
			split()
			for _, v := range *values {
				if !yield(v) {
					return
				}
			}
		}}
	}
	return half(&s.in), half(&s.out)
}

// Partition returns the pairs for which match returns true, and those for
// which it returns false, with their keys. It reads the collection once,
// when either result is first read, and keeps both, as
// [Collection.Partition] does.
func (c KeyedCollection[K, V]) Partition(match func(K, V) bool) (KeyedCollection[K, V], KeyedCollection[K, V]) {
	s := &partitionState[Pair[K, V]]{}
	split := func() {
		s.split(func(add func(Pair[K, V], bool)) {
			for k, v := range c.Seq() {
				add(Pair[K, V]{k, v}, match(k, v))
			}
		})
	}
	half := func(pairs *[]Pair[K, V]) KeyedCollection[K, V] {
		return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
			split()
			for _, p := range *pairs {
				if !yield(p.First, p.Second) {
					return
				}
			}
		}}
	}
	return half(&s.in), half(&s.out)
}

// partitionState keeps the two halves once a read of the source ends.
type partitionState[E any] struct {
	mu      sync.Mutex
	done    bool
	in, out []E
}

// split runs read, which adds each value to a half, unless a read has
// ended before. The halves are kept only when read returns: a panic keeps
// nothing, so that the next read starts again.
func (s *partitionState[E]) split(read func(add func(E, bool))) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.done {
		return
	}
	var in, out []E
	read(func(e E, matched bool) {
		if matched {
			in = append(in, e)
		} else {
			out = append(out, e)
		}
	})
	s.in, s.out, s.done = in, out, true
}

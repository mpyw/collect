package collect

import "sync"

// Partition returns the values for which match returns true, and those for
// which it returns false. It reads the collection once, when either result
// is first read, and keeps both, as Laravel's partition reads once. Later
// reads of either return the kept values.
func (c Collection[T]) Partition(match func(T) bool) (Collection[T], Collection[T]) {
	var (
		once     sync.Once
		in, out  []T
		readOnce = func() {
			once.Do(func() {
				for v := range c.Seq() {
					if match(v) {
						in = append(in, v)
					} else {
						out = append(out, v)
					}
				}
			})
		}
	)
	half := func(values *[]T) Collection[T] {
		return Collection[T]{seq: func(yield func(T) bool) {
			readOnce()
			for _, v := range *values {
				if !yield(v) {
					return
				}
			}
		}}
	}
	return half(&in), half(&out)
}

// Partition returns the pairs for which match returns true, and those for
// which it returns false, with their keys. It reads the collection once,
// when either result is first read, and keeps both.
func (c KeyedCollection[K, V]) Partition(match func(K, V) bool) (KeyedCollection[K, V], KeyedCollection[K, V]) {
	type half struct {
		ks []K
		vs []V
	}
	var (
		once    sync.Once
		in, out half
	)
	readOnce := func() {
		once.Do(func() {
			for k, v := range c.Seq() {
				h := &out
				if match(k, v) {
					h = &in
				}
				h.ks, h.vs = append(h.ks, k), append(h.vs, v)
			}
		})
	}
	read := func(h *half) KeyedCollection[K, V] {
		return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
			readOnce()
			for i, k := range h.ks {
				if !yield(k, h.vs[i]) {
					return
				}
			}
		}}
	}
	return read(&in), read(&out)
}

package collect

import "slices"

// Reverse returns the values last first. It reads every value before the
// first.
func (c Collection[T]) Reverse() Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		all := c.All()
		for _, v := range slices.Backward(all) {
			if !yield(v) {
				return
			}
		}
	}}
}

// Reverse returns the pairs last first, with their keys. It reads every
// pair before the first.
func (c KeyedCollection[K, V]) Reverse() KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		var (
			ks []K
			vs []V
		)
		for k, v := range c.Seq() {
			ks, vs = append(ks, k), append(vs, v)
		}
		for i := len(ks) - 1; i >= 0; i-- {
			if !yield(ks[i], vs[i]) {
				return
			}
		}
	}}
}

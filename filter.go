package collect

// Filter returns a collection of the values for which keep returns true.
func (c Collection[T]) Filter(keep func(T) bool) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		for v := range c.Seq() {
			if keep(v) && !yield(v) {
				return
			}
		}
	}}
}

// Filter returns the pairs for which keep returns true, with their keys.
func (c KeyedCollection[K, V]) Filter(keep func(K, V) bool) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for k, v := range c.Seq() {
			if keep(k, v) && !yield(k, v) {
				return
			}
		}
	}}
}

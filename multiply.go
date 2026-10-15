package collect

// Multiply returns the values n times over. It reads the collection n
// times. A negative n is 0.
func (c Collection[T]) Multiply(n int) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		for range max(n, 0) {
			for v := range c.Seq() {
				if !yield(v) {
					return
				}
			}
		}
	}}
}

// Multiply returns the pairs n times over, with their keys. It reads the
// collection n times. A negative n is 0.
func (c KeyedCollection[K, V]) Multiply(n int) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for range max(n, 0) {
			for k, v := range c.Seq() {
				if !yield(k, v) {
					return
				}
			}
		}
	}}
}

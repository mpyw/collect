package collect

// Skip returns a collection of the values after the first n. A negative n
// is 0.
func (c Collection[T]) Skip(n int) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		i := 0
		for v := range c.Seq() {
			if i < n {
				i++
				continue
			}
			if !yield(v) {
				return
			}
		}
	}}
}

// Skip returns the pairs after the first n. A negative n is 0.
func (c KeyedCollection[K, V]) Skip(n int) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		i := 0
		for k, v := range c.Seq() {
			if i < n {
				i++
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}}
}

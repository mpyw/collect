package collect

// Map returns a collection of f applied to each value.
func (c Collection[T]) Map[U any](f func(T) U) Collection[U] {
	return Collection[U]{seq: func(yield func(U) bool) {
		for v := range c.Seq() {
			if !yield(f(v)) {
				return
			}
		}
	}}
}

// Map returns the keys with f applied to each pair as their values.
func (c KeyedCollection[K, V]) Map[U any](f func(K, V) U) KeyedCollection[K, U] {
	return KeyedCollection[K, U]{seq: func(yield func(K, U) bool) {
		for k, v := range c.Seq() {
			if !yield(k, f(k, v)) {
				return
			}
		}
	}}
}

package collect

// MapWithKeys returns the key and value f returns for each value.
func (c Collection[T]) MapWithKeys[K comparable, U any](f func(T) (K, U)) KeyedCollection[K, U] {
	return KeyedCollection[K, U]{seq: func(yield func(K, U) bool) {
		for v := range c.Seq() {
			if !yield(f(v)) {
				return
			}
		}
	}}
}

// MapWithKeys returns the key and value f returns for each pair.
func (c KeyedCollection[K, V]) MapWithKeys[K2 comparable, U any](f func(K, V) (K2, U)) KeyedCollection[K2, U] {
	return KeyedCollection[K2, U]{seq: func(yield func(K2, U) bool) {
		for k, v := range c.Seq() {
			if !yield(f(k, v)) {
				return
			}
		}
	}}
}

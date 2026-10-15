package collect

// TakeUntil returns the values before the first for which match returns
// true. It reads no value past that one.
func (c Collection[T]) TakeUntil(match func(T) bool) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		for v := range c.Seq() {
			if match(v) || !yield(v) {
				return
			}
		}
	}}
}

// TakeUntil returns the pairs before the first for which match returns
// true. It reads no pair past that one.
func (c KeyedCollection[K, V]) TakeUntil(match func(K, V) bool) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for k, v := range c.Seq() {
			if match(k, v) || !yield(k, v) {
				return
			}
		}
	}}
}

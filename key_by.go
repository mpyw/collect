package collect

// KeyBy returns the values with key(v) as the key of each v.
func (c Collection[T]) KeyBy[K comparable](key func(T) K) KeyedCollection[K, T] {
	return KeyedCollection[K, T]{seq: func(yield func(K, T) bool) {
		for v := range c.Seq() {
			if !yield(key(v), v) {
				return
			}
		}
	}}
}

// KeyBy returns the values with the key f returns for each pair as their
// new key.
func (c KeyedCollection[K, V]) KeyBy[K2 comparable](f func(K, V) K2) KeyedCollection[K2, V] {
	return KeyedCollection[K2, V]{seq: func(yield func(K2, V) bool) {
		for k, v := range c.Seq() {
			if !yield(f(k, v), v) {
				return
			}
		}
	}}
}

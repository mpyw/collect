package collect

// Reject returns a collection of the values for which drop returns false.
func (c Collection[T]) Reject(drop func(T) bool) Collection[T] {
	return c.Filter(func(v T) bool { return !drop(v) })
}

// Reject returns the pairs for which drop returns false, with their keys.
func (c KeyedCollection[K, V]) Reject(drop func(K, V) bool) KeyedCollection[K, V] {
	return c.Filter(func(k K, v V) bool { return !drop(k, v) })
}

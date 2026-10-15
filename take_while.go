package collect

// TakeWhile returns the values up to the first for which match returns
// false. It reads no value past that one.
func (c Collection[T]) TakeWhile(match func(T) bool) Collection[T] {
	return c.TakeUntil(func(v T) bool { return !match(v) })
}

// TakeWhile returns the pairs up to the first for which match returns
// false. It reads no pair past that one.
func (c KeyedCollection[K, V]) TakeWhile(match func(K, V) bool) KeyedCollection[K, V] {
	return c.TakeUntil(func(k K, v V) bool { return !match(k, v) })
}

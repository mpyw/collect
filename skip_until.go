package collect

// SkipUntil returns the values from the first for which match returns true.
func (c Collection[T]) SkipUntil(match func(T) bool) Collection[T] {
	return c.SkipWhile(func(v T) bool { return !match(v) })
}

// SkipUntil returns the pairs from the first for which match returns true.
func (c KeyedCollection[K, V]) SkipUntil(match func(K, V) bool) KeyedCollection[K, V] {
	return c.SkipWhile(func(k K, v V) bool { return !match(k, v) })
}

package collect

// Every reports whether match returns true for every value. It is true for
// an empty collection. It reads no value past the first that fails.
func (c Collection[T]) Every(match func(T) bool) bool {
	return !c.Contains(func(v T) bool { return !match(v) })
}

// Every reports whether match returns true for every pair. It is true for
// an empty collection. It reads no pair past the first that fails.
func (c KeyedCollection[K, V]) Every(match func(K, V) bool) bool {
	return !c.Contains(func(k K, v V) bool { return !match(k, v) })
}

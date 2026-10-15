package collect

// FirstWhere returns the first value for which match returns true. It is
// false if there is none. It reads no value past that one.
func (c Collection[T]) FirstWhere(match func(T) bool) (T, bool) {
	return c.Filter(match).First()
}

// FirstWhere returns the first pair for which match returns true. It is
// false if there is none. It reads no pair past that one.
func (c KeyedCollection[K, V]) FirstWhere(match func(K, V) bool) (K, V, bool) {
	return c.Filter(match).First()
}

package collect

// MustFirst returns the first value. It panics with [ErrItemNotFound] if
// there is none, as Laravel's firstOrFail throws. It reads at most one.
// [Collection.First] reports the absence instead.
func (c Collection[T]) MustFirst() T {
	v, ok := c.First()
	if !ok {
		panic(ErrItemNotFound)
	}
	return v
}

// MustFirst returns the first pair. It panics with [ErrItemNotFound] if
// there is none. It reads at most one.
func (c KeyedCollection[K, V]) MustFirst() (K, V) {
	k, v, ok := c.First()
	if !ok {
		panic(ErrItemNotFound)
	}
	return k, v
}

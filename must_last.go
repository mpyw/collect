package collect

// MustLast returns the last value. It panics with [ErrItemNotFound] if
// there is none. It reads every value. [Collection.Last] reports the
// absence instead.
func (c Collection[T]) MustLast() T {
	v, ok := c.Last()
	if !ok {
		panic(ErrItemNotFound)
	}
	return v
}

// MustLast returns the last pair. It panics with [ErrItemNotFound] if
// there is none. It reads every pair.
func (c KeyedCollection[K, V]) MustLast() (K, V) {
	k, v, ok := c.Last()
	if !ok {
		panic(ErrItemNotFound)
	}
	return k, v
}

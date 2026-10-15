package collect

// MustSole returns the only value. It panics with the error
// [Collection.Sole] would return, [ErrItemNotFound] or
// [ErrMultipleItemsFound], as Laravel's sole throws. It reads at most two.
func (c Collection[T]) MustSole() T {
	v, err := c.Sole()
	if err != nil {
		panic(err)
	}
	return v
}

// MustSole returns the only pair. It panics with the error
// [KeyedCollection.Sole] would return.
func (c KeyedCollection[K, V]) MustSole() (K, V) {
	k, v, err := c.Sole()
	if err != nil {
		panic(err)
	}
	return k, v
}

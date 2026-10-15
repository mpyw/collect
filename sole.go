package collect

// Sole returns the only value. The error is [ErrItemNotFound] if there is
// none, and [ErrMultipleItemsFound] if there is more than one. It reads at
// most two. Filter first for the only value that matches.
func (c Collection[T]) Sole() (T, error) {
	var zero T
	values := c.Take(2).All()
	switch len(values) {
	case 0:
		return zero, ErrItemNotFound
	case 1:
		return values[0], nil
	default:
		return zero, ErrMultipleItemsFound
	}
}

// Sole returns the only pair. The error is [ErrItemNotFound] if there is
// none, and [ErrMultipleItemsFound] if there is more than one. It reads at
// most two.
func (c KeyedCollection[K, V]) Sole() (K, V, error) {
	var (
		zeroK K
		zeroV V
	)
	ks, vs := pairsOf(c.Take(2))
	switch len(ks) {
	case 0:
		return zeroK, zeroV, ErrItemNotFound
	case 1:
		return ks[0], vs[0], nil
	default:
		return zeroK, zeroV, ErrMultipleItemsFound
	}
}

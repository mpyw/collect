package collect

// Search returns the index, from 0, of the first value for which match
// returns true. It is false if there is none.
func (c Collection[T]) Search(match func(T) bool) (int, bool) {
	k, _, ok := c.Keyed().FirstWhere(func(_ int, v T) bool { return match(v) })
	return k, ok
}

// Search returns the key of the first pair for which match returns true.
// It is false if there is none.
func (c KeyedCollection[K, V]) Search(match func(K, V) bool) (K, bool) {
	k, _, ok := c.FirstWhere(match)
	return k, ok
}

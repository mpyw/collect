package collect

// Contains reports whether match returns true for a value. It reads no
// value past that one. To look for a value, compare in match.
func (c Collection[T]) Contains(match func(T) bool) bool {
	_, ok := c.FirstWhere(match)
	return ok
}

// Contains reports whether match returns true for a pair. It reads no
// pair past that one.
func (c KeyedCollection[K, V]) Contains(match func(K, V) bool) bool {
	_, _, ok := c.FirstWhere(match)
	return ok
}

package collect

// DoesntContain reports whether match returns false for every value.
func (c Collection[T]) DoesntContain(match func(T) bool) bool {
	return !c.Contains(match)
}

// DoesntContain reports whether match returns false for every pair.
func (c KeyedCollection[K, V]) DoesntContain(match func(K, V) bool) bool {
	return !c.Contains(match)
}

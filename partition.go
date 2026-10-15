package collect

// Partition returns the values for which match returns true, and those for
// which it returns false. Each reads the collection on its own.
func (c Collection[T]) Partition(match func(T) bool) (Collection[T], Collection[T]) {
	return c.Filter(match), c.Reject(match)
}

// Partition returns the pairs for which match returns true, and those for
// which it returns false. Each reads the collection on its own.
func (c KeyedCollection[K, V]) Partition(match func(K, V) bool) (KeyedCollection[K, V], KeyedCollection[K, V]) {
	return c.Filter(match), c.Reject(match)
}

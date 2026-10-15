package collect

// ForPage returns the values of page page, from 1, with perPage values to a
// page. A page before the first is the first.
func (c Collection[T]) ForPage(page, perPage int) Collection[T] {
	return c.Skip(max(page-1, 0) * perPage).Take(perPage)
}

// ForPage returns the pairs of page page, from 1, with perPage pairs to a
// page, with their keys.
func (c KeyedCollection[K, V]) ForPage(page, perPage int) KeyedCollection[K, V] {
	return c.Skip(max(page-1, 0) * perPage).Take(perPage)
}

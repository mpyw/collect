package collect

// HasSole reports whether the collection has exactly one value. It reads at
// most two. Filter first to ask about the values that match.
func (c Collection[T]) HasSole() bool {
	return c.Take(2).Count() == 1
}

// HasSole reports whether the collection has exactly one pair. It reads at
// most two.
func (c KeyedCollection[K, V]) HasSole() bool {
	return c.Take(2).Count() == 1
}

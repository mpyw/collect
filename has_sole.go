package collect

// HasSole reports whether the collection has exactly one value. It reads at
// most two, and drops them, as [Collection.IsEmpty] does. Filter first to
// ask about the values that match. [Collection.Sole] returns the value too.
func (c Collection[T]) HasSole() bool {
	return c.Take(2).Count() == 1
}

// HasSole reports whether the collection has exactly one pair. It reads at
// most two, and drops them, as [Collection.IsEmpty] does.
func (c KeyedCollection[K, V]) HasSole() bool {
	return c.Take(2).Count() == 1
}

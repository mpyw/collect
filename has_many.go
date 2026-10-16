package collect

// HasMany reports whether the collection has more than one value. It reads
// at most two, and drops them, as [Collection.IsEmpty] does.
func (c Collection[T]) HasMany() bool {
	return c.Take(2).Count() == 2
}

// HasMany reports whether the collection has more than one pair. It reads
// at most two, and drops them, as [Collection.IsEmpty] does.
func (c KeyedCollection[K, V]) HasMany() bool {
	return c.Take(2).Count() == 2
}

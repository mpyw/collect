package collect

// IsNotEmpty reports whether the collection has a value. It reads at most
// one, and drops it, as [Collection.IsEmpty] does.
// A source that can be read once loses what it reads: call Remember first.
func (c Collection[T]) IsNotEmpty() bool {
	return !c.IsEmpty()
}

// IsNotEmpty reports whether the collection has a pair. It reads at most
// one, and drops it, as [Collection.IsEmpty] does.
// A source that can be read once loses what it reads: call Remember first.
func (c KeyedCollection[K, V]) IsNotEmpty() bool {
	return !c.IsEmpty()
}

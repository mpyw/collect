package collect

// IsNotEmpty reports whether the collection has a value. It reads at most
// one.
func (c Collection[T]) IsNotEmpty() bool {
	return !c.IsEmpty()
}

// IsNotEmpty reports whether the collection has a pair. It reads at most
// one.
func (c KeyedCollection[K, V]) IsNotEmpty() bool {
	return !c.IsEmpty()
}

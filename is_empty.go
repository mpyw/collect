package collect

// IsEmpty reports whether the collection has no values. It reads at most
// one.
func (c Collection[T]) IsEmpty() bool {
	for range c.Seq() {
		return false
	}
	return true
}

// IsEmpty reports whether the collection has no pairs. It reads at most
// one.
func (c KeyedCollection[K, V]) IsEmpty() bool {
	for range c.Seq() {
		return false
	}
	return true
}

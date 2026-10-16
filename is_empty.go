package collect

// IsEmpty reports whether the collection has no values. It reads at most
// one.
//
// It drops the values it reads. A source that can be read once loses them:
// call Remember first. A source that runs a query on each read runs it
// again when the collection is read after.
func (c Collection[T]) IsEmpty() bool {
	for range c.Seq() {
		return false
	}
	return true
}

// IsEmpty reports whether the collection has no pairs. It reads at most
// one, as [Collection.IsEmpty] does, and drops it.
// A source that can be read once loses what it reads: call Remember first.
func (c KeyedCollection[K, V]) IsEmpty() bool {
	for range c.Seq() {
		return false
	}
	return true
}

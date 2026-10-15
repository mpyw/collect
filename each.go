package collect

// Each calls f with each value, in order. To stop early, range over
// [Collection.Seq] and break.
func (c Collection[T]) Each(f func(T)) {
	for v := range c.Seq() {
		f(v)
	}
}

// Each calls f with each pair, in order. To stop early, range over
// [KeyedCollection.Seq] and break.
func (c KeyedCollection[K, V]) Each(f func(K, V)) {
	for k, v := range c.Seq() {
		f(k, v)
	}
}

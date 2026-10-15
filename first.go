package collect

// First returns the first value. It is false if there is none. It reads
// at most one.
func (c Collection[T]) First() (T, bool) {
	for v := range c.Seq() {
		return v, true
	}
	var zero T
	return zero, false
}

// First returns the first pair. It is false if there is none. It reads at
// most one.
func (c KeyedCollection[K, V]) First() (K, V, bool) {
	for k, v := range c.Seq() {
		return k, v, true
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV, false
}

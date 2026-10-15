package collect

// Last returns the last value. It is false if there is none. It reads
// every value.
func (c Collection[T]) Last() (T, bool) {
	var (
		last T
		ok   bool
	)
	for v := range c.Seq() {
		last, ok = v, true
	}
	return last, ok
}

// Last returns the last pair. It is false if there is none. It reads every
// pair.
func (c KeyedCollection[K, V]) Last() (K, V, bool) {
	var (
		lastK K
		lastV V
		ok    bool
	)
	for k, v := range c.Seq() {
		lastK, lastV, ok = k, v, true
	}
	return lastK, lastV, ok
}

package collect

// After returns the value just after the first for which match returns
// true. It is false if no value matches, or if the last value is the first
// that does.
func (c Collection[T]) After(match func(T) bool) (T, bool) {
	found := false
	for v := range c.Seq() {
		if found {
			return v, true
		}
		found = match(v)
	}
	var zero T
	return zero, false
}

// After returns the pair just after the first for which match returns
// true. It is false if no pair matches, or if the last pair is the first
// that does.
func (c KeyedCollection[K, V]) After(match func(K, V) bool) (K, V, bool) {
	found := false
	for k, v := range c.Seq() {
		if found {
			return k, v, true
		}
		found = match(k, v)
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV, false
}

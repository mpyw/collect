package collect

// Before returns the value just before the first for which match returns
// true. It is false if no value matches, or if the first value does.
func (c Collection[T]) Before(match func(T) bool) (T, bool) {
	var (
		previous T
		has      bool
	)
	for v := range c.Seq() {
		if match(v) {
			return previous, has
		}
		previous, has = v, true
	}
	var zero T
	return zero, false
}

// Before returns the pair just before the first for which match returns
// true. It is false if no pair matches, or if the first pair does.
func (c KeyedCollection[K, V]) Before(match func(K, V) bool) (K, V, bool) {
	var (
		prevK K
		prevV V
		has   bool
	)
	for k, v := range c.Seq() {
		if match(k, v) {
			return prevK, prevV, has
		}
		prevK, prevV, has = k, v, true
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV, false
}

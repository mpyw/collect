package collect

// DiffAssocBy returns the pairs that other does not have: other has no such
// key, or f returns something else for its value of the key. Pass
// [Identity] to compare the values. When other has a key twice, its last
// value counts. It reads other whole before the first pair.
func (c KeyedCollection[K, V]) DiffAssocBy[E comparable](other KeyedCollection[K, V], f func(V) E) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		theirs := map[K]E{}
		for k, v := range other.Seq() {
			theirs[k] = f(v)
		}
		for k, v := range c.Seq() {
			if e, ok := theirs[k]; ok && e == f(v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}}
}

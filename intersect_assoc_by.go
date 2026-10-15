package collect

// IntersectAssocBy returns the pairs that other has too: the same key, with a
// value of the same key, as f returns it. Pass [Identity] to compare the
// values. When other has a key twice, its last value counts. It reads other
// whole before the first pair.
func (c KeyedCollection[K, V]) IntersectAssocBy[E comparable](other KeyedCollection[K, V], f func(V) E) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		theirs := map[K]E{}
		for k, v := range other.Seq() {
			theirs[k] = f(v)
		}
		for k, v := range c.Seq() {
			if e, ok := theirs[k]; ok && e == f(v) && !yield(k, v) {
				return
			}
		}
	}}
}

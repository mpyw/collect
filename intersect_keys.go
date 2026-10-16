package collect

// IntersectKeys returns the pairs whose key is the key of a pair of
// other. It is Laravel's intersectByKeys, named as [KeyedCollection.DiffKeys]
// is. It reads other whole before the first pair.
func (c KeyedCollection[K, V]) IntersectKeys(other KeyedCollection[K, V]) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		in := map[K]struct{}{}
		for k := range other.Seq() {
			in[k] = struct{}{}
		}
		for k, v := range c.Seq() {
			if _, ok := in[k]; ok && !yield(k, v) {
				return
			}
		}
	}}
}

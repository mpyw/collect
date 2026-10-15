package collect

// Union returns the pairs, then the pairs of other whose key is not the key
// of a pair before, as PHP's + on arrays does.
func (c KeyedCollection[K, V]) Union(other KeyedCollection[K, V]) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		seen := map[K]struct{}{}
		for k, v := range c.Seq() {
			seen[k] = struct{}{}
			if !yield(k, v) {
				return
			}
		}
		for k, v := range other.Seq() {
			if _, ok := seen[k]; ok {
				continue
			}
			seen[k] = struct{}{}
			if !yield(k, v) {
				return
			}
		}
	}}
}

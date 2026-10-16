package collect

// Union returns the pairs, then the pairs of other whose key is not the key
// of a pair before, as PHP's + on arrays does. When other has a key twice,
// its last value counts, as Laravel reads other into an array. It reads
// other whole before the first pair.
func (c KeyedCollection[K, V]) Union(other KeyedCollection[K, V]) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		theirs := map[K]V{}
		var order []K
		for k, v := range other.Seq() {
			if _, ok := theirs[k]; !ok {
				order = append(order, k)
			}
			theirs[k] = v
		}
		seen := map[K]struct{}{}
		for k, v := range c.Seq() {
			seen[k] = struct{}{}
			if !yield(k, v) {
				return
			}
		}
		for _, k := range order {
			if _, ok := seen[k]; ok {
				continue
			}
			if !yield(k, theirs[k]) {
				return
			}
		}
	}}
}

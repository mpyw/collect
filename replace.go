package collect

// Replace returns the pairs, with the value of other for each key that
// other has, then the pairs of other whose key comes in neither before. It
// reads other whole before the first pair. It is also Laravel's merge of
// keyed collections: a Go key is never renumbered.
func (c KeyedCollection[K, V]) Replace(other KeyedCollection[K, V]) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		values := map[K]V{}
		var order []K
		for k, v := range other.Seq() {
			if _, ok := values[k]; !ok {
				order = append(order, k)
			}
			values[k] = v
		}
		for k, v := range c.Seq() {
			if w, ok := values[k]; ok {
				v = w
				delete(values, k)
			}
			if !yield(k, v) {
				return
			}
		}
		for _, k := range order {
			if v, ok := values[k]; ok && !yield(k, v) {
				return
			}
		}
	}}
}

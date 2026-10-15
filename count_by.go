package collect

// CountBy counts the values by the key f returns for each. The counts are
// in the order their keys first come. It reads every value before it
// yields the first count. Pass [Identity] to count equal values.
func (c Collection[T]) CountBy[K comparable](f func(T) K) KeyedCollection[K, int] {
	return KeyedCollection[K, int]{seq: func(yield func(K, int) bool) {
		counts := map[K]int{}
		var order []K
		for v := range c.Seq() {
			k := f(v)
			if counts[k] == 0 {
				order = append(order, k)
			}
			counts[k]++
		}
		for _, k := range order {
			if !yield(k, counts[k]) {
				return
			}
		}
	}}
}

// CountBy counts the pairs by the key f returns for each. The counts are in
// the order their keys first come. It reads every pair before it yields the
// first count.
func (c KeyedCollection[K, V]) CountBy[G comparable](f func(K, V) G) KeyedCollection[G, int] {
	return KeyedCollection[G, int]{seq: func(yield func(G, int) bool) {
		var gs []G
		for k, v := range c.Seq() {
			gs = append(gs, f(k, v))
		}
		for g, n := range FromSlice(gs).CountBy(Identity).Seq() {
			if !yield(g, n) {
				return
			}
		}
	}}
}

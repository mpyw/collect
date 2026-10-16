package collect

// ModeBy returns the keys that f returns most often for the values, in the
// order they first come. It is empty for an empty collection. It reads
// every value before it yields the first key. Pass [Identity] for a
// collection of keys.
func (c Collection[T]) ModeBy[K comparable](f func(T) K) Collection[K] {
	return Collection[K]{seq: func(yield func(K) bool) {
		keys, groups := groupsOf(func(add func(K, struct{})) {
			for v := range c.Seq() {
				add(f(v), struct{}{})
			}
		})
		modeByYield(keys, groups, yield)
	}}
}

// ModeBy returns the keys that f returns most often for the pairs, in the
// order they first come. It is empty for an empty collection. It reads
// every pair before it yields the first key.
func (c KeyedCollection[K, V]) ModeBy[E comparable](f func(K, V) E) Collection[E] {
	return Collection[E]{seq: func(yield func(E) bool) {
		keys, groups := groupsOf(func(add func(E, struct{})) {
			for k, v := range c.Seq() {
				add(f(k, v), struct{}{})
			}
		})
		modeByYield(keys, groups, yield)
	}}
}

// modeByYield yields the keys whose groups are the largest, in order.
func modeByYield[K any](keys []K, groups [][]struct{}, yield func(K) bool) {
	most := 0
	for _, g := range groups {
		most = max(most, len(g))
	}
	for i, k := range keys {
		if len(groups[i]) == most && !yield(k) {
			return
		}
	}
}

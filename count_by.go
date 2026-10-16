package collect

// CountBy counts the values by the key f returns for each. The counts are
// in the order their keys first come. It reads every value before it
// yields the first count. Pass [Identity] to count equal values.
func (c Collection[T]) CountBy[K comparable](f func(T) K) KeyedCollection[K, int] {
	return KeyedCollection[K, int]{seq: func(yield func(K, int) bool) {
		keys, groups := groupsOf(func(add func(K, struct{})) {
			for v := range c.Seq() {
				add(f(v), struct{}{})
			}
		})
		for i, k := range keys {
			if !yield(k, len(groups[i])) {
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
		keys, groups := groupsOf(func(add func(G, struct{})) {
			for k, v := range c.Seq() {
				add(f(k, v), struct{}{})
			}
		})
		for i, g := range keys {
			if !yield(g, len(groups[i])) {
				return
			}
		}
	}}
}

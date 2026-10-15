package collect

// FlipBy returns the pairs with the key f returns for each value as the key,
// and the old key as the value. Pass [Identity] when the values are
// comparable.
func (c KeyedCollection[K, V]) FlipBy[K2 comparable](f func(V) K2) KeyedCollection[K2, K] {
	return KeyedCollection[K2, K]{seq: func(yield func(K2, K) bool) {
		for k, v := range c.Seq() {
			if !yield(f(v), k) {
				return
			}
		}
	}}
}

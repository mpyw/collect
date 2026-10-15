package collect

// Keys returns the keys without their values. A key that comes more than
// once is there each time.
func (c KeyedCollection[K, V]) Keys() Collection[K] {
	return Collection[K]{seq: func(yield func(K) bool) {
		for k := range c.Seq() {
			if !yield(k) {
				return
			}
		}
	}}
}

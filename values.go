package collect

// Values returns the values without their keys, as [maps.Values] does.
func (c KeyedCollection[K, V]) Values() Collection[V] {
	return Collection[V]{seq: func(yield func(V) bool) {
		for _, v := range c.Seq() {
			if !yield(v) {
				return
			}
		}
	}}
}

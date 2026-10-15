package collect

// Eager reads every value now, and returns a collection of them. The new
// collection can be read any number of times, as one from [Of] can.
func (c Collection[T]) Eager() Collection[T] {
	return FromSlice(c.All())
}

// Eager reads every pair now, and returns a collection of them, in order
// and with every key, even one that comes more than once.
func (c KeyedCollection[K, V]) Eager() KeyedCollection[K, V] {
	var (
		ks []K
		vs []V
	)
	for k, v := range c.Seq() {
		ks, vs = append(ks, k), append(vs, v)
	}
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for i, k := range ks {
			if !yield(k, vs[i]) {
				return
			}
		}
	}}
}

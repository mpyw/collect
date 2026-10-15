package collect

// IntersectBy returns the values whose key, as f returns it, is the key of a
// value of other. Pass [Identity] to compare the values. It reads other
// whole before the first value.
func (c Collection[T]) IntersectBy[K comparable](other Collection[T], f func(T) K) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		in := map[K]struct{}{}
		for o := range other.Seq() {
			in[f(o)] = struct{}{}
		}
		for v := range c.Seq() {
			if _, ok := in[f(v)]; ok && !yield(v) {
				return
			}
		}
	}}
}

// IntersectBy returns the pairs whose value has a key, as f returns it, that
// a value of other has. Each pair keeps its own key. It reads other whole
// before the first pair.
func (c KeyedCollection[K, V]) IntersectBy[E comparable](other KeyedCollection[K, V], f func(V) E) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		in := map[E]struct{}{}
		for _, o := range other.Seq() {
			in[f(o)] = struct{}{}
		}
		for k, v := range c.Seq() {
			if _, ok := in[f(v)]; ok && !yield(k, v) {
				return
			}
		}
	}}
}

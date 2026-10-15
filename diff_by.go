package collect

// DiffBy returns the values whose key, as f returns it, is not the key of a
// value of other. Pass [Identity] to compare the values. It reads other
// whole before the first value.
func (c Collection[T]) DiffBy[K comparable](other Collection[T], f func(T) K) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		in := diffByKeySet(other, f)
		for v := range c.Seq() {
			if _, ok := in[f(v)]; !ok && !yield(v) {
				return
			}
		}
	}}
}

func diffByKeySet[T any, K comparable](c Collection[T], f func(T) K) map[K]struct{} {
	set := map[K]struct{}{}
	for v := range c.Seq() {
		set[f(v)] = struct{}{}
	}
	return set
}

// DiffBy returns the pairs whose value has a key, as f returns it, that no
// value of other has. Each pair keeps its own key, as in Laravel. It reads
// other whole before the first pair.
func (c KeyedCollection[K, V]) DiffBy[E comparable](other KeyedCollection[K, V], f func(V) E) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		in := diffByKeySet(other.Values(), f)
		for k, v := range c.Seq() {
			if _, ok := in[f(v)]; !ok && !yield(k, v) {
				return
			}
		}
	}}
}

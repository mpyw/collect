package collect

// IntersectFunc returns the values that equal a value of other, as eq
// tells. It compares each value with each of other, which it reads whole
// first.
func (c Collection[T]) IntersectFunc(other Collection[T], eq func(a, b T) bool) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		others := other.All()
		for v := range c.Seq() {
			if FromSlice(others).Contains(func(o T) bool { return eq(v, o) }) && !yield(v) {
				return
			}
		}
	}}
}

// IntersectFunc returns the pairs whose value equals a value of other, as
// eq tells. Each pair keeps its own key.
func (c KeyedCollection[K, V]) IntersectFunc(other KeyedCollection[K, V], eq func(a, b V) bool) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		others := other.Values().All()
		for k, v := range c.Seq() {
			if FromSlice(others).Contains(func(o V) bool { return eq(v, o) }) && !yield(k, v) {
				return
			}
		}
	}}
}

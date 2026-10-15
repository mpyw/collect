package collect

// Multiply returns the values n times over. Each read reads the collection
// once, as Laravel's multiply does, and repeats what it read. A negative n
// is 0.
func (c Collection[T]) Multiply(n int) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		if n <= 0 {
			return
		}
		var read []T
		for v := range c.Seq() {
			if !yield(v) {
				return
			}
			read = append(read, v)
		}
		for range n - 1 {
			for _, v := range read {
				if !yield(v) {
					return
				}
			}
		}
	}}
}

// Multiply returns the pairs n times over, with their keys. Each read reads
// the collection once, and repeats what it read. A negative n is 0.
func (c KeyedCollection[K, V]) Multiply(n int) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		if n <= 0 {
			return
		}
		var (
			ks []K
			vs []V
		)
		for k, v := range c.Seq() {
			if !yield(k, v) {
				return
			}
			ks, vs = append(ks, k), append(vs, v)
		}
		for range n - 1 {
			for i, k := range ks {
				if !yield(k, vs[i]) {
					return
				}
			}
		}
	}}
}

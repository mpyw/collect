package collect

// UniqueBy returns the first value of each key, as f returns it. Pass
// [Identity] to drop equal values. It keeps every key it has seen.
func (c Collection[T]) UniqueBy[K comparable](f func(T) K) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		seen := map[K]struct{}{}
		for v := range c.Seq() {
			k := f(v)
			if _, ok := seen[k]; ok {
				continue
			}
			seen[k] = struct{}{}
			if !yield(v) {
				return
			}
		}
	}}
}

// UniqueBy returns the first pair of each key, as f returns it, with its own
// key.
func (c KeyedCollection[K, V]) UniqueBy[E comparable](f func(K, V) E) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		seen := map[E]struct{}{}
		for k, v := range c.Seq() {
			e := f(k, v)
			if _, ok := seen[e]; ok {
				continue
			}
			seen[e] = struct{}{}
			if !yield(k, v) {
				return
			}
		}
	}}
}

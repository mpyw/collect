package collect

// DuplicatesBy returns what f returns for each value, each time it returns
// what it returned for a value before, as Laravel's duplicates does. Pass
// [Identity] for the repeated values themselves.
func (c Collection[T]) DuplicatesBy[K comparable](f func(T) K) Collection[K] {
	return Collection[K]{seq: func(yield func(K) bool) {
		seen := map[K]struct{}{}
		for v := range c.Seq() {
			k := f(v)
			if _, ok := seen[k]; !ok {
				seen[k] = struct{}{}
				continue
			}
			if !yield(k) {
				return
			}
		}
	}}
}

// DuplicatesBy returns what f returns for each pair, each time it returns
// what it returned for a pair before, with that pair's own key, as
// Laravel's duplicates does.
func (c KeyedCollection[K, V]) DuplicatesBy[E comparable](f func(K, V) E) KeyedCollection[K, E] {
	return KeyedCollection[K, E]{seq: func(yield func(K, E) bool) {
		seen := map[E]struct{}{}
		for k, v := range c.Seq() {
			e := f(k, v)
			if _, ok := seen[e]; !ok {
				seen[e] = struct{}{}
				continue
			}
			if !yield(k, e) {
				return
			}
		}
	}}
}

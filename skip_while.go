package collect

// SkipWhile returns the values from the first for which match returns
// false.
func (c Collection[T]) SkipWhile(match func(T) bool) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		skipping := true
		for v := range c.Seq() {
			if skipping = skipping && match(v); !skipping && !yield(v) {
				return
			}
		}
	}}
}

// SkipWhile returns the pairs from the first for which match returns false.
func (c KeyedCollection[K, V]) SkipWhile(match func(K, V) bool) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		skipping := true
		for k, v := range c.Seq() {
			if skipping = skipping && match(k, v); !skipping && !yield(k, v) {
				return
			}
		}
	}}
}

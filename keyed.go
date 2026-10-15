package collect

// Keyed returns the values with their indexes, from 0, as keys.
func (c Collection[T]) Keyed() KeyedCollection[int, T] {
	return KeyedCollection[int, T]{seq: func(yield func(int, T) bool) {
		i := 0
		for v := range c.Seq() {
			if !yield(i, v) {
				return
			}
			i++
		}
	}}
}

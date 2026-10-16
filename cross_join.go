package collect

// CrossJoin returns a pair of each value with each value of other, in
// order. It reads other whole before the first pair. P is always
// Entry[T, U], and Go infers it, as with [Collection.Chunk].
func (c Collection[T]) CrossJoin[U any, E interface{ Entry[T, U] }](other Collection[U]) Collection[E] {
	return Collection[E]{seq: func(yield func(E) bool) {
		others := other.All()
		for t := range c.Seq() {
			for _, u := range others {
				if !yield(E(Entry[T, U]{t, u})) {
					return
				}
			}
		}
	}}
}

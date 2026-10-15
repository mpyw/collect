package collect

// CrossJoin returns a pair of each value with each value of other, in
// order. It reads other whole before the first pair. P is always
// Pair[T, U], and Go infers it, as with [Collection.Chunk].
func (c Collection[T]) CrossJoin[U any, P interface{ Pair[T, U] }](other Collection[U]) Collection[P] {
	return Collection[P]{seq: func(yield func(P) bool) {
		others := other.All()
		for t := range c.Seq() {
			for _, u := range others {
				if !yield(P(Pair[T, U]{t, u})) {
					return
				}
			}
		}
	}}
}

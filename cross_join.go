package collect

// CrossJoin returns each value with each value of other, in order, as a
// [Tuple2]. It reads other whole before the first one. An empty other
// gives none, and reads none of the collection. Z is always
// Tuple2[T, U], and Go infers it, as with [Collection.Chunk].
func (c Collection[T]) CrossJoin[U any, Z interface{ Tuple2[T, U] }](other Collection[U]) Collection[Z] {
	return Collection[Z]{seq: func(yield func(Z) bool) {
		others := other.All()
		if len(others) == 0 {
			// Ranging over an endless collection would yield nothing, and
			// never end.
			return
		}
		for t := range c.Seq() {
			for _, u := range others {
				if !yield(Z(Tuple2[T, U]{t, u})) {
					return
				}
			}
		}
	}}
}

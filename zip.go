package collect

import "iter"

// Zip returns a pair of a value and the value of other at the same place.
// It goes on to the end of the longer one, as Laravel does: the missing
// side of a pair is the zero value. P is always Pair[T, U], and Go infers
// it, as with [Collection.Chunk].
func (c Collection[T]) Zip[U any, P interface{ Pair[T, U] }](other Collection[U]) Collection[P] {
	return Collection[P]{seq: func(yield func(P) bool) {
		nextT, stopT := iter.Pull(c.Seq())
		defer stopT()
		nextU, stopU := iter.Pull(other.Seq())
		defer stopU()
		for {
			t, okT := nextT()
			u, okU := nextU()
			if !okT && !okU {
				return
			}
			if !yield(P(Pair[T, U]{t, u})) {
				return
			}
		}
	}}
}

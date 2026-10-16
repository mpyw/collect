package collect

import "iter"

// Zip returns each value with the value of other at the same place, as a
// [Tuple2]: A is the value, and B is other's. It goes on to the end of the
// longer one, as Laravel does: the missing side is the zero value. Z is
// always Tuple2[T, U], and Go infers it, as with [Collection.Chunk].
// [Collection.Zip3], [Collection.Zip4] and [Collection.Zip5] zip more.
func (c Collection[T]) Zip[U any, Z interface{ Tuple2[T, U] }](other Collection[U]) Collection[Z] {
	return Collection[Z]{seq: func(yield func(Z) bool) {
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
			if !yield(Z(Tuple2[T, U]{t, u})) {
				return
			}
		}
	}}
}

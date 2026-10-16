package collect

import "iter"

// Zip returns each value with the value of other at the same place, as an
// [Entry]: Key is the value, and Value is other's. It goes on to the end of
// the longer one, as Laravel does: the missing side is the zero value. E is
// always Entry[T, U], and Go infers it, as with [Collection.Chunk].
// [Collection.Zip3], [Collection.Zip4] and [Collection.Zip5] zip more.
func (c Collection[T]) Zip[U any, E interface{ Entry[T, U] }](other Collection[U]) Collection[E] {
	return Collection[E]{seq: func(yield func(E) bool) {
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
			if !yield(E(Entry[T, U]{t, u})) {
				return
			}
		}
	}}
}

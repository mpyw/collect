package collect

import "iter"

// Zip3 returns each value with the values of us and vs at the same place,
// as a [Tuple3], as [Collection.Zip] does for two collections. It goes on
// to the end of the longest of the three: a missing value is the zero
// value. Z is always Tuple3[T, U, V], and Go infers it.
func (c Collection[T]) Zip3[U, V any, Z interface{ Tuple3[T, U, V] }](us Collection[U], vs Collection[V]) Collection[Z] {
	return Collection[Z]{seq: func(yield func(Z) bool) {
		nextT, stopT := iter.Pull(c.Seq())
		defer stopT()
		nextU, stopU := iter.Pull(us.Seq())
		defer stopU()
		nextV, stopV := iter.Pull(vs.Seq())
		defer stopV()
		for {
			t, okT := nextT()
			u, okU := nextU()
			v, okV := nextV()
			if !okT && !okU && !okV {
				return
			}
			if !yield(Z(Tuple3[T, U, V]{t, u, v})) {
				return
			}
		}
	}}
}

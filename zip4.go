package collect

import "iter"

// Zip4 returns each value with the values of us, vs and ws at the same
// place, as a [Tuple4], as [Collection.Zip2] does for two collections. It
// goes on to the end of the longest of the four: a missing value is the
// zero value. Z is always Tuple4[T, U, V, W], and Go infers it.
func (c Collection[T]) Zip4[U, V, W any, Z interface{ Tuple4[T, U, V, W] }](us Collection[U], vs Collection[V], ws Collection[W]) Collection[Z] {
	return Collection[Z]{seq: func(yield func(Z) bool) {
		nextT, stopT := iter.Pull(c.Seq())
		defer stopT()
		nextU, stopU := iter.Pull(us.Seq())
		defer stopU()
		nextV, stopV := iter.Pull(vs.Seq())
		defer stopV()
		nextW, stopW := iter.Pull(ws.Seq())
		defer stopW()
		for {
			t, okT := nextT()
			u, okU := nextU()
			v, okV := nextV()
			w, okW := nextW()
			if !okT && !okU && !okV && !okW {
				return
			}
			if !yield(Z(Tuple4[T, U, V, W]{t, u, v, w})) {
				return
			}
		}
	}}
}

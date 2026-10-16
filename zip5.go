package collect

import "iter"

// Zip5 returns each value with the values of us, vs, ws and xs at the same
// place, as a [Tuple5], as [Collection.Zip] does for two collections. It
// goes on to the end of the longest of the five: a missing value is the
// zero value. Z is always Tuple5[T, U, V, W, X], and Go infers it.
func (c Collection[T]) Zip5[U, V, W, X any, Z interface{ Tuple5[T, U, V, W, X] }](us Collection[U], vs Collection[V], ws Collection[W], xs Collection[X]) Collection[Z] {
	return Collection[Z]{seq: func(yield func(Z) bool) {
		nextT, stopT := iter.Pull(c.Seq())
		defer stopT()
		nextU, stopU := iter.Pull(us.Seq())
		defer stopU()
		nextV, stopV := iter.Pull(vs.Seq())
		defer stopV()
		nextW, stopW := iter.Pull(ws.Seq())
		defer stopW()
		nextX, stopX := iter.Pull(xs.Seq())
		defer stopX()
		for {
			t, okT := nextT()
			u, okU := nextU()
			v, okV := nextV()
			w, okW := nextW()
			x, okX := nextX()
			if !okT && !okU && !okV && !okW && !okX {
				return
			}
			if !yield(Z(Tuple5[T, U, V, W, X]{t, u, v, w, x})) {
				return
			}
		}
	}}
}

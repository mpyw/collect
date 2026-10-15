package collect

import "slices"

// ChunkWhile returns runs of values. A run goes on while keep returns true
// for the next value and the run so far. The run is valid only during the
// call. C is always Collection[T], and Go infers it, as with
// [Collection.Chunk].
func (c Collection[T]) ChunkWhile[C interface{ Collection[T] }](keep func(v T, chunk Collection[T]) bool) Collection[C] {
	return Collection[C]{seq: func(yield func(C) bool) {
		var buf []T
		for v := range c.Seq() {
			if len(buf) > 0 && !keep(v, FromSlice(buf)) {
				if !yield(C(Of(buf...))) {
					return
				}
				buf = buf[:0]
			}
			buf = append(buf, v)
		}
		if len(buf) > 0 {
			yield(C(Of(buf...)))
		}
	}}
}

// ChunkWhile returns runs of pairs. A run goes on while keep returns true
// for the next pair and the run so far. Each pair keeps its own key. The
// run is valid only during the call. C is always KeyedCollection[K, V], and
// Go infers it.
func (c KeyedCollection[K, V]) ChunkWhile[C interface{ KeyedCollection[K, V] }](keep func(k K, v V, chunk KeyedCollection[K, V]) bool) Collection[C] {
	return Collection[C]{seq: func(yield func(C) bool) {
		var (
			ks []K
			vs []V
		)
		for k, v := range c.Seq() {
			if len(ks) > 0 && !keep(k, v, keyedOf(ks, vs)) {
				if !yield(C(keyedOf(slices.Clone(ks), slices.Clone(vs)))) {
					return
				}
				ks, vs = ks[:0], vs[:0]
			}
			ks, vs = append(ks, k), append(vs, v)
		}
		if len(ks) > 0 {
			yield(C(keyedOf(ks, vs)))
		}
	}}
}

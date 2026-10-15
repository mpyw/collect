package collect

import "slices"

// ChunkBy returns runs of adjacent values that have the same key, as f
// returns it. C is always Collection[T], and Go infers it.
func (c Collection[T]) ChunkBy[K comparable, C interface{ Collection[T] }](f func(T) K) Collection[C] {
	return Collection[C]{seq: func(yield func(C) bool) {
		var (
			buf  []T
			last K
		)
		for v := range c.Seq() {
			k := f(v)
			if len(buf) > 0 && k != last {
				if !yield(C(Of(buf...))) {
					return
				}
				buf = buf[:0]
			}
			buf, last = append(buf, v), k
		}
		if len(buf) > 0 {
			yield(C(Of(buf...)))
		}
	}}
}

// ChunkBy returns runs of adjacent pairs that have the same key, as f
// returns it. Each pair keeps its own key. C is always KeyedCollection[K, V],
// and Go infers it.
func (c KeyedCollection[K, V]) ChunkBy[E comparable, C interface{ KeyedCollection[K, V] }](f func(K, V) E) Collection[C] {
	return Collection[C]{seq: func(yield func(C) bool) {
		var (
			ks   []K
			vs   []V
			last E
		)
		for k, v := range c.Seq() {
			e := f(k, v)
			if len(ks) > 0 && e != last {
				if !yield(C(keyedOf(slices.Clone(ks), slices.Clone(vs)))) {
					return
				}
				ks, vs = ks[:0], vs[:0]
			}
			ks, vs, last = append(ks, k), append(vs, v), e
		}
		if len(ks) > 0 {
			yield(C(keyedOf(ks, vs)))
		}
	}}
}

package collect

import "slices"

// Chunk returns the values in collections of size values each. The last
// one has fewer if the values run out. It panics if size is less than 1,
// as [slices.Chunk] does.
//
// C is always Collection[T], and Go infers it: write c.Chunk(2). See
// [Collection] for why it is a type parameter. The result is a Collection
// like any other, so it chunks again: c.Chunk(2).Chunk(3).
func (c Collection[T]) Chunk[C interface{ Collection[T] }](size int) Collection[C] {
	chunkCheckSize(size)
	return Collection[C]{seq: func(yield func(C) bool) {
		// Each chunk gets a slice of its own: a chunk may be kept, and
		// read, after the next one is made. Of copies buf.
		buf := make([]T, 0, size)
		for v := range c.Seq() {
			if buf = append(buf, v); len(buf) < size {
				continue
			}
			if !yield(C(Of(buf...))) {
				return
			}
			buf = buf[:0]
		}
		if len(buf) > 0 {
			yield(C(Of(buf...)))
		}
	}}
}

// Chunk returns the pairs in keyed collections of size pairs each, as
// Laravel's chunk does: each pair keeps its key. The last one has fewer if
// the pairs run out. It panics if size is less than 1.
//
// C is always KeyedCollection[K, V], and Go infers it, as with
// [Collection.Chunk].
func (c KeyedCollection[K, V]) Chunk[C interface{ KeyedCollection[K, V] }](size int) Collection[C] {
	chunkCheckSize(size)
	return Collection[C]{seq: func(yield func(C) bool) {
		ks := make([]K, 0, size)
		vs := make([]V, 0, size)
		for k, v := range c.Seq() {
			if ks, vs = append(ks, k), append(vs, v); len(ks) < size {
				continue
			}
			if !yield(C(keyedOf(slices.Clone(ks), slices.Clone(vs)))) {
				return
			}
			ks, vs = ks[:0], vs[:0]
		}
		if len(ks) > 0 {
			yield(C(keyedOf(slices.Clone(ks), slices.Clone(vs))))
		}
	}}
}

func chunkCheckSize(size int) {
	if size < 1 {
		panic("collect: Chunk size must be at least 1")
	}
}

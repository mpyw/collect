package collect

// SplitIn returns the values in n groups, each full but the last, as
// Laravel's splitIn does. It reads every value first, and panics with [ErrInvalidArgument] if n is
// less than 1. C is always Collection[T], and Go infers it.
func (c Collection[T]) SplitIn[C interface{ Collection[T] }](n int) Collection[C] {
	if n < 1 {
		invalidArgument("SplitIn needs at least 1 group, not %d", n)
	}
	return Collection[C]{seq: func(yield func(C) bool) {
		// Not c.Chunk: a method that calls a pinned method of its own
		// receiver type is an instantiation cycle.
		all := c.All()
		size := splitInSize(len(all), n)
		for start := 0; start < len(all); start += size {
			if !yield(C(Of(all[start:min(start+size, len(all))]...))) {
				return
			}
		}
	}}
}

// SplitIn returns the pairs in n groups, each full but the last, as
// [Collection.SplitIn] does. Each pair keeps its own key. C is always
// KeyedCollection[K, V], and Go infers it.
func (c KeyedCollection[K, V]) SplitIn[C interface{ KeyedCollection[K, V] }](n int) Collection[C] {
	if n < 1 {
		invalidArgument("SplitIn needs at least 1 group, not %d", n)
	}
	return Collection[C]{seq: func(yield func(C) bool) {
		ks, vs := pairsOf(c)
		size := splitInSize(len(ks), n)
		for start := 0; start < len(ks); start += size {
			end := min(start+size, len(ks))
			if !yield(C(keyedOf(ks[start:end], vs[start:end]))) {
				return
			}
		}
	}}
}

// splitInSize returns the size of each group of count values in n groups,
// rounded up, with no overflow for a large n.
func splitInSize(count, n int) int {
	size := count / n
	if count%n != 0 {
		size++
	}
	return size
}

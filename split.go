package collect

// Split returns the values in n groups, as even in size as can be. The
// first groups have one more value when they cannot all be the same. There
// are fewer than n groups if there are fewer than n values. It reads every
// value first, and panics with [ErrInvalidArgument] if n is less than 1. C is always Collection[T],
// and Go infers it.
func (c Collection[T]) Split[C interface{ Collection[T] }](n int) Collection[C] {
	if n < 1 {
		invalidArgument("Split needs at least 1 group, not %d", n)
	}
	return Collection[C]{seq: func(yield func(C) bool) {
		all := c.All()
		size, extra := len(all)/n, len(all)%n
		start := 0
		for i := range n {
			end := start + size
			if i < extra {
				end++
			}
			if end == start {
				return
			}
			if !yield(C(Of(all[start:end]...))) {
				return
			}
			start = end
		}
	}}
}

// Split returns the pairs in n groups, as even in size as can be, as
// [Collection.Split] does. Each pair keeps its own key. C is always
// KeyedCollection[K, V], and Go infers it.
func (c KeyedCollection[K, V]) Split[C interface{ KeyedCollection[K, V] }](n int) Collection[C] {
	if n < 1 {
		invalidArgument("Split needs at least 1 group, not %d", n)
	}
	return Collection[C]{seq: func(yield func(C) bool) {
		ks, vs := pairsOf(c)
		size, extra := len(ks)/n, len(ks)%n
		start := 0
		for i := range n {
			end := start + size
			if i < extra {
				end++
			}
			if end == start || !yield(C(keyedOf(ks[start:end], vs[start:end]))) {
				return
			}
			start = end
		}
	}}
}

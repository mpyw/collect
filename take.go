package collect

import "math"

// Take returns a collection of the first n values, or of all of them if
// there are fewer. It reads no value past the nth, so it ends an endless
// sequence.
//
// A negative n takes the last -n values instead, as Laravel does. Those
// are known only at the end, so it reads every value, and keeps -n of them.
func (c Collection[T]) Take(n int) Collection[T] {
	if n < 0 {
		return Collection[T]{seq: func(yield func(T) bool) {
			for _, v := range takeLast(c.Seq(), n) {
				if !yield(v) {
					return
				}
			}
		}}
	}
	return Collection[T]{seq: func(yield func(T) bool) {
		if n == 0 {
			return
		}
		i := 0
		for v := range c.Seq() {
			if !yield(v) {
				return
			}
			if i++; i == n {
				return
			}
		}
	}}
}

// Take returns the first n pairs, or all of them if there are fewer. It
// reads no pair past the nth. A negative n takes the last -n pairs, and
// reads every pair.
func (c KeyedCollection[K, V]) Take(n int) KeyedCollection[K, V] {
	if n < 0 {
		return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
			// takeLast reads to the end, so yield never returns false.
			pairs := func(yield func(Pair[K, V]) bool) {
				for k, v := range c.Seq() {
					yield(Pair[K, V]{k, v})
				}
			}
			for _, p := range takeLast(pairs, n) {
				if !yield(p.First, p.Second) {
					return
				}
			}
		}}
	}
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		if n == 0 {
			return
		}
		i := 0
		for k, v := range c.Seq() {
			if !yield(k, v) {
				return
			}
			if i++; i == n {
				return
			}
		}
	}}
}

// takeLast returns the last -negative values of seq, in order, keeping no
// more than that many at a time. math.MinInt, whose negation overflows,
// takes them all.
func takeLast[E any](seq func(func(E) bool), negative int) []E {
	n := -negative
	if n < 0 {
		n = math.MaxInt
	}
	// The ring grows as values come, so a count far past the length of the
	// source allocates nothing up front.
	var ring []E
	next := 0
	for v := range seq {
		if len(ring) < n {
			ring = append(ring, v)
			continue
		}
		ring[next] = v
		next = (next + 1) % n
	}
	return append(ring[next:], ring[:next]...)
}

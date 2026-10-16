package collect

import "math"

// Pad returns the values, then copies of v until there are size values. A
// negative size puts the copies before the values, until there are -size,
// and reads every value first.
func (c Collection[T]) Pad(size int, v T) Collection[T] {
	if size < 0 {
		return Collection[T]{seq: func(yield func(T) bool) {
			all := c.All()
			target := -size
			if target < 0 {
				target = math.MaxInt // -math.MinInt overflows.
			}
			for range target - len(all) {
				if !yield(v) {
					return
				}
			}
			for _, w := range all {
				if !yield(w) {
					return
				}
			}
		}}
	}
	return Collection[T]{seq: func(yield func(T) bool) {
		n := 0
		for w := range c.Seq() {
			if !yield(w) {
				return
			}
			n++
		}
		for ; n < size; n++ {
			if !yield(v) {
				return
			}
		}
	}}
}

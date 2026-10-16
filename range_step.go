package collect

// RangeStep returns the numbers from from to to, both included, step apart,
// as Laravel's range does: it adds step while the next number is not past
// to. It counts down if to is less than from, whatever the sign of step.
//
// It stops before a number that would wrap around the end of N, or that
// would not change, as with a float step too small for from. A NaN in any
// argument gives no numbers. It panics with [ErrInvalidArgument] if step is
// 0, as Laravel throws.
func RangeStep[N Number](from, to, step N) Collection[N] {
	if step == 0 {
		invalidArgument("RangeStep step must not be 0")
	}
	if step < 0 {
		step = -step
	}
	return Collection[N]{seq: func(yield func(N) bool) {
		// A NaN is the only value that is not equal to itself.
		if from != from || to != to || step != step {
			return
		}
		up := from <= to
		for n := from; yield(n); {
			next := n - step
			if up {
				next = n + step
			}
			if up && (next <= n || next > to) || !up && (next >= n || next < to) {
				return
			}
			n = next
		}
	}}
}

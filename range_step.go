package collect

// RangeStep returns the numbers from from to to, both included, step apart.
// It counts down if to is less than from, whatever the sign of step.
//
// An integer range adds step to each number. A float range computes each
// number as from plus i times step, as PHP's range does, so that rounding
// does not build up: RangeStep(0, 1, 0.1) ends at 1. Laravel's lazy range
// adds the step instead, and ends at 0.9999999999999999.
//
// It stops before a number that would wrap around the end of N, or that
// would not change, as with a float step too small for from. A NaN in any
// argument gives no numbers. It panics with an error that wraps
// [ErrInvalidArgument] if step is 0, as Laravel throws.
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
		var half N = 1
		if half /= 2; half != 0 {
			rangeStepFloat(from, to, step, up, yield)
			return
		}
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

// rangeStepFloat yields from plus or minus i times step, for i from 0,
// while the number does not pass to and still changes.
func rangeStepFloat[N Number](from, to, step N, up bool, yield func(N) bool) {
	last := from
	for i := 0; ; i++ {
		// The conversion rounds the product on its own, as PHP does. Go may
		// otherwise fuse the multiply and the add into one instruction, and
		// round only once.
		offset := N(N(i) * step)
		n := from - offset
		if up {
			n = from + offset
		}
		if i > 0 && (n == last || up && n > to || !up && n < to) {
			return
		}
		if !yield(n) {
			return
		}
		last = n
	}
}

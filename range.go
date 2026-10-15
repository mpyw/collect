package collect

// Range returns the numbers from from to to, both included, a step of 1
// apart. It counts down if to is less than from.
func Range[N Number](from, to N) Collection[N] {
	return RangeStep(from, to, 1)
}

// RangeStep returns the numbers from from to to, both included, step apart.
// It counts down if to is less than from, whatever the sign of step. It
// panics if step is 0, as Laravel throws.
func RangeStep[N Number](from, to, step N) Collection[N] {
	if step == 0 {
		panic("collect: RangeStep step must not be 0")
	}
	if step < 0 {
		step = -step
	}
	return Collection[N]{seq: func(yield func(N) bool) {
		if from <= to {
			// Compare before adding, so that a step past the top of N
			// does not wrap around.
			for n := from; ; n += step {
				if !yield(n) || to-n < step {
					return
				}
			}
		}
		for n := from; ; n -= step {
			if !yield(n) || n-to < step {
				return
			}
		}
	}}
}

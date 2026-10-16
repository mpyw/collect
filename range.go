package collect

// Range returns the numbers from from to to, both included, a step of 1
// apart. It counts down if to is less than from.
func Range[N Number](from, to N) Collection[N] {
	return RangeStep(from, to, 1)
}

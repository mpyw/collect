package collect

import "slices"

// MedianBy returns the middle of the numbers f returns for the values, or the
// mean of the two middle ones. It is false for an empty collection. It
// keeps every number. Pass [Identity] for a collection of numbers.
func (c Collection[T]) MedianBy[N Number](f func(T) N) (float64, bool) {
	numbers := slices.Sorted(c.Map(f).Seq())
	n := len(numbers)
	switch {
	case n == 0:
		return 0, false
	case n%2 == 1:
		return float64(numbers[n/2]), true
	default:
		return (float64(numbers[n/2-1]) + float64(numbers[n/2])) / 2, true
	}
}

// MedianBy returns the middle of the numbers f returns for the pairs, or the
// mean of the two middle ones. It is false for an empty collection.
func (c KeyedCollection[K, V]) MedianBy[N Number](f func(K, V) N) (float64, bool) {
	var numbers []N
	for k, v := range c.Seq() {
		numbers = append(numbers, f(k, v))
	}
	return FromSlice(numbers).MedianBy(Identity)
}

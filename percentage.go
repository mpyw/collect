package collect

import "math"

// Percentage returns the percentage of the values for which match returns
// true, rounded to precision decimal places. It is false for an empty
// collection.
func (c Collection[T]) Percentage(match func(T) bool, precision int) (float64, bool) {
	var all, matched int
	for v := range c.Seq() {
		all++
		if match(v) {
			matched++
		}
	}
	if all == 0 {
		return 0, false
	}
	return percentageRound(float64(matched)*100/float64(all), precision), true
}

// Percentage returns the percentage of the pairs for which match returns
// true, rounded to precision decimal places. It is false for an empty
// collection.
func (c KeyedCollection[K, V]) Percentage(match func(K, V) bool, precision int) (float64, bool) {
	var matched []bool
	for k, v := range c.Seq() {
		matched = append(matched, match(k, v))
	}
	return FromSlice(matched).Percentage(Identity, precision)
}

// percentageRound rounds p, from 0 to 100, to precision decimal places. A
// float64 holds no more than 15 of them, and 100 rounds to 0 from 3 places
// to the left, so precision is clamped to that range: math.Pow10 of a
// larger one is an infinity, and the result would be NaN.
func percentageRound(p float64, precision int) float64 {
	scale := math.Pow10(min(max(precision, -3), 15))
	return math.Round(p*scale) / scale
}

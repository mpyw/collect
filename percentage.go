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
	scale := math.Pow10(precision)
	return math.Round(float64(matched)*100/float64(all)*scale) / scale, true
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

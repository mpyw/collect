package collect

import "math"

// ForPage returns the values of page page, from 1, with perPage values to a
// page. A page before the first is the first. It is Laravel's forPage:
// Slice(max(0, (page-1)*perPage), perPage), so a negative perPage counts
// from the end, as [Collection.Slice] does.
func (c Collection[T]) ForPage(page, perPage int) Collection[T] {
	return c.Slice(forPageOffset(page, perPage), perPage)
}

// ForPage returns the pairs of page page, from 1, with perPage pairs to a
// page, with their keys, as [Collection.ForPage] does.
func (c KeyedCollection[K, V]) ForPage(page, perPage int) KeyedCollection[K, V] {
	return c.Slice(forPageOffset(page, perPage), perPage)
}

// forPageOffset returns max(0, (page-1)*perPage), as Laravel computes it,
// with no overflow: a product past math.MaxInt is math.MaxInt. Both
// factors negative make it positive.
func forPageOffset(page, perPage int) int {
	if page == 1 || perPage == 0 {
		return 0
	}
	if page < 1 != (perPage < 0) {
		return 0 // a negative product
	}
	// Both factors have the same sign now. Count in magnitudes, with page-1
	// and math.MinInt, whose negations overflow, as math.MaxInt.
	a, b := page-1, perPage
	if page < 1 {
		a, b = -a, -b
		if page == math.MinInt || a < 0 {
			a = math.MaxInt
		}
		if b < 0 {
			b = math.MaxInt
		}
	}
	if a > math.MaxInt/b {
		return math.MaxInt
	}
	return a * b
}

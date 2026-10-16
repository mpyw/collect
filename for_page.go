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

// forPageOffset returns max(0, (page-1)*perPage), with no overflow: a
// product past math.MaxInt is math.MaxInt.
func forPageOffset(page, perPage int) int {
	if page <= 1 || perPage <= 0 {
		return 0
	}
	if page-1 > math.MaxInt/perPage {
		return math.MaxInt
	}
	return (page - 1) * perPage
}

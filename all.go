package collect

import "slices"

// All reads the collection into a new slice. It is Laravel's all. Note that
// [slices.All] is an iterator: here, [Collection.Seq] is.
func (c Collection[T]) All() []T {
	return slices.Collect(c.Seq())
}

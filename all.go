package collect

import (
	"maps"
	"slices"
)

// All reads the collection into a new slice.
func (c Collection[T]) All() []T {
	return slices.Collect(c.Seq())
}

// All reads the collection into a new map. When a key comes more than
// once, its last value is kept.
func (c KeyedCollection[K, V]) All() map[K]V {
	return maps.Collect(c.Seq())
}

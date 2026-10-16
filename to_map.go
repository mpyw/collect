package collect

import "maps"

// ToMap reads every pair into a new map. A map has no order, and keeps one
// value of each key: when a key comes more than once, its last value wins.
// Entries().All() keeps the order and every pair. It stands for Laravel's
// all, which keeps the order of a PHP array.
func (c KeyedCollection[K, V]) ToMap() map[K]V {
	return maps.Collect(c.Seq())
}

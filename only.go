package collect

import "slices"

// Only returns the first pair of each of keys, in the collection's order.
// It reads no pair past the one that completes them, as Laravel's lazy
// only does.
func (c KeyedCollection[K, V]) Only(keys ...K) KeyedCollection[K, V] {
	// Copied now, as Except does: a change to keys after the call changes
	// nothing.
	keys = slices.Clone(keys)
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		wanted := make(map[K]struct{}, len(keys))
		for _, k := range keys {
			wanted[k] = struct{}{}
		}
		if len(wanted) == 0 {
			return
		}
		for k, v := range c.Seq() {
			if _, ok := wanted[k]; !ok {
				continue
			}
			if !yield(k, v) {
				return
			}
			if delete(wanted, k); len(wanted) == 0 {
				return
			}
		}
	}}
}

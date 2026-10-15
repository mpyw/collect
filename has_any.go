package collect

// HasAny reports whether any of keys is the key of a pair. It reads no pair
// past the first that has one.
func (c KeyedCollection[K, V]) HasAny(keys ...K) bool {
	return c.Contains(func(k K, _ V) bool {
		for _, want := range keys {
			if k == want {
				return true
			}
		}
		return false
	})
}

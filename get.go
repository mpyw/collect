package collect

// Get returns the value of the first pair with key k. It is false if there
// is none. It reads no pair past that one. Note that [KeyedCollection.ToMap]
// keeps the last value of a key instead.
func (c KeyedCollection[K, V]) Get(k K) (V, bool) {
	_, v, ok := c.FirstWhere(func(key K, _ V) bool { return key == k })
	return v, ok
}

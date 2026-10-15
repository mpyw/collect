package collect

// Has reports whether every one of keys is the key of a pair. It reads no
// pair past the one that completes them. It is true for no keys.
func (c KeyedCollection[K, V]) Has(keys ...K) bool {
	missing := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		missing[k] = struct{}{}
	}
	if len(missing) == 0 {
		return true
	}
	for k := range c.Seq() {
		if delete(missing, k); len(missing) == 0 {
			return true
		}
	}
	return false
}

package collect

// Except returns the pairs whose key is none of keys.
func (c KeyedCollection[K, V]) Except(keys ...K) KeyedCollection[K, V] {
	dropped := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		dropped[k] = struct{}{}
	}
	return c.Reject(func(k K, _ V) bool {
		_, ok := dropped[k]
		return ok
	})
}

package collect

// WhereNotInBy returns the values whose key, as f returns it, is none of
// keys.
func (c Collection[T]) WhereNotInBy[K comparable](f func(T) K, keys ...K) Collection[T] {
	set := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		set[k] = struct{}{}
	}
	return c.Reject(func(v T) bool {
		_, ok := set[f(v)]
		return ok
	})
}

// WhereNotInBy returns the pairs whose key, as f returns it, is none of keys.
func (c KeyedCollection[K, V]) WhereNotInBy[E comparable](f func(K, V) E, keys ...E) KeyedCollection[K, V] {
	set := make(map[E]struct{}, len(keys))
	for _, k := range keys {
		set[k] = struct{}{}
	}
	return c.Reject(func(k K, v V) bool {
		_, ok := set[f(k, v)]
		return ok
	})
}

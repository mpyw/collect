package collect

// WhereInBy returns the values whose key, as f returns it, is one of keys.
func (c Collection[T]) WhereInBy[K comparable](f func(T) K, keys ...K) Collection[T] {
	set := whereInBySet(keys)
	return c.Filter(func(v T) bool {
		_, ok := set[f(v)]
		return ok
	})
}

func whereInBySet[K comparable](keys []K) map[K]struct{} {
	set := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		set[k] = struct{}{}
	}
	return set
}

// WhereInBy returns the pairs whose key, as f returns it, is one of keys.
func (c KeyedCollection[K, V]) WhereInBy[E comparable](f func(K, V) E, keys ...E) KeyedCollection[K, V] {
	set := whereInBySet(keys)
	return c.Filter(func(k K, v V) bool {
		_, ok := set[f(k, v)]
		return ok
	})
}

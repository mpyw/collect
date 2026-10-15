package collect

// TapEach returns the same values, calling f with each as it is read.
func (c Collection[T]) TapEach(f func(T)) Collection[T] {
	return c.Map(func(v T) T {
		f(v)
		return v
	})
}

// TapEach returns the same pairs, calling f with each as it is read.
func (c KeyedCollection[K, V]) TapEach(f func(K, V)) KeyedCollection[K, V] {
	return c.Map(func(k K, v V) V {
		f(k, v)
		return v
	})
}

package collect

// When returns f applied to the collection if cond is true, and the
// collection itself if not. It keeps a chain going through a condition.
func (c Collection[T]) When(cond bool, f func(Collection[T]) Collection[T]) Collection[T] {
	if cond {
		return f(c)
	}
	return c
}

// When returns f applied to the collection if cond is true, and the
// collection itself if not.
func (c KeyedCollection[K, V]) When(cond bool, f func(KeyedCollection[K, V]) KeyedCollection[K, V]) KeyedCollection[K, V] {
	if cond {
		return f(c)
	}
	return c
}

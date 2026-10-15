package collect

// Unless returns f applied to the collection if cond is false, and the
// collection itself if not.
func (c Collection[T]) Unless(cond bool, f func(Collection[T]) Collection[T]) Collection[T] {
	return c.When(!cond, f)
}

// Unless returns f applied to the collection if cond is false, and the
// collection itself if not.
func (c KeyedCollection[K, V]) Unless(cond bool, f func(KeyedCollection[K, V]) KeyedCollection[K, V]) KeyedCollection[K, V] {
	return c.When(!cond, f)
}

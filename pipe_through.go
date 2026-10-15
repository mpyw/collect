package collect

// PipeThrough passes the collection through each of fs in turn, and
// returns what the last returns.
func (c Collection[T]) PipeThrough(fs ...func(Collection[T]) Collection[T]) Collection[T] {
	for _, f := range fs {
		c = f(c)
	}
	return c
}

// PipeThrough passes the collection through each of fs in turn, and
// returns what the last returns.
func (c KeyedCollection[K, V]) PipeThrough(fs ...func(KeyedCollection[K, V]) KeyedCollection[K, V]) KeyedCollection[K, V] {
	for _, f := range fs {
		c = f(c)
	}
	return c
}

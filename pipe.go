package collect

// Pipe returns what f returns for the collection.
func (c Collection[T]) Pipe[U any](f func(Collection[T]) U) U {
	return f(c)
}

// Pipe returns what f returns for the collection.
func (c KeyedCollection[K, V]) Pipe[U any](f func(KeyedCollection[K, V]) U) U {
	return f(c)
}

package collect

// MustEnsure returns the values as U, lazily. Write U: c.MustEnsure[*Dog]().
// Reading a value that is not a U panics with the error [Collection.Ensure]
// would return, which wraps [ErrUnexpectedValue], as Laravel's ensure
// throws.
func (c Collection[T]) MustEnsure[U any]() Collection[U] {
	return Collection[U]{seq: func(yield func(U) bool) {
		for v := range c.Seq() {
			u, err := asType[T, U](v)
			if err != nil {
				panic(err)
			}
			if !yield(u) {
				return
			}
		}
	}}
}

// MustEnsure returns the pairs with their values as U, lazily. Reading a
// value that is not a U panics with the error [KeyedCollection.Ensure]
// would return.
func (c KeyedCollection[K, V]) MustEnsure[U any]() KeyedCollection[K, U] {
	return KeyedCollection[K, U]{seq: func(yield func(K, U) bool) {
		for k, v := range c.Seq() {
			u, err := asType[V, U](v)
			if err != nil {
				panic(err)
			}
			if !yield(k, u) {
				return
			}
		}
	}}
}

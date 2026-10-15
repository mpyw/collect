package collect

// FlatMap returns the values of the collections f returns, in order, as one
// collection. With a function that returns its argument, it flattens a
// collection of collections, such as the result of [Collection.Chunk].
func (c Collection[T]) FlatMap[U any](f func(T) Collection[U]) Collection[U] {
	return Collection[U]{seq: func(yield func(U) bool) {
		for v := range c.Seq() {
			for u := range f(v).Seq() {
				if !yield(u) {
					return
				}
			}
		}
	}}
}

// FlatMap returns the pairs of the keyed collections f returns, in order,
// as one keyed collection.
func (c KeyedCollection[K, V]) FlatMap[K2 comparable, U any](f func(K, V) KeyedCollection[K2, U]) KeyedCollection[K2, U] {
	return KeyedCollection[K2, U]{seq: func(yield func(K2, U) bool) {
		for k, v := range c.Seq() {
			for k2, u := range f(k, v).Seq() {
				if !yield(k2, u) {
					return
				}
			}
		}
	}}
}

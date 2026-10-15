package collect

// Concat returns the values of the collection, then of each of others. It
// is also Laravel's merge of lists.
func (c Collection[T]) Concat(others ...Collection[T]) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		for _, source := range append([]Collection[T]{c}, others...) {
			for v := range source.Seq() {
				if !yield(v) {
					return
				}
			}
		}
	}}
}

// Concat returns the pairs of the collection, then of each of others, with
// their keys. A key may then come more than once.
func (c KeyedCollection[K, V]) Concat(others ...KeyedCollection[K, V]) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for _, source := range append([]KeyedCollection[K, V]{c}, others...) {
			for k, v := range source.Seq() {
				if !yield(k, v) {
					return
				}
			}
		}
	}}
}

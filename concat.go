package collect

import "slices"

// Concat returns the values of the collection, then of each of others. It
// is also Laravel's merge of lists.
func (c Collection[T]) Concat(others ...Collection[T]) Collection[T] {
	// A new slice: a change to others after the call changes nothing.
	sources := slices.Concat([]Collection[T]{c}, others)
	return Collection[T]{seq: func(yield func(T) bool) {
		for _, source := range sources {
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
	sources := slices.Concat([]KeyedCollection[K, V]{c}, others)
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for _, source := range sources {
			for k, v := range source.Seq() {
				if !yield(k, v) {
					return
				}
			}
		}
	}}
}

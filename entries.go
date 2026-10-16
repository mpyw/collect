package collect

// Entries returns each value with its index, from 0, as an [Entry], as
// ECMAScript's Array.prototype.entries does. E is always Entry[int, T], and
// Go infers it, as with [Collection.Chunk].
func (c Collection[T]) Entries[E interface{ Entry[int, T] }]() Collection[E] {
	return Collection[E]{seq: func(yield func(E) bool) {
		i := 0
		for v := range c.Seq() {
			if !yield(E(Entry[int, T]{i, v})) {
				return
			}
			i++
		}
	}}
}

// Entries returns each pair as an [Entry], in order and with every repeated
// key, as ECMAScript's Map.prototype.entries does. Entries().All() is the
// pairs in order, which [KeyedCollection.ToMap] cannot keep. E is always
// Entry[K, V], and Go infers it, as with [Collection.Chunk].
func (c KeyedCollection[K, V]) Entries[E interface{ Entry[K, V] }]() Collection[E] {
	return Collection[E]{seq: func(yield func(E) bool) {
		for k, v := range c.Seq() {
			if !yield(E(Entry[K, V]{k, v})) {
				return
			}
		}
	}}
}

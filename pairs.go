package collect

// Pairs returns each pair of c as a [Pair], in order: First is the key, and
// Second the value. It keeps the order and every repeated key, which
// [KeyedCollection.All] does not.
//
// It is a function, not a method: a method of KeyedCollection that returns
// a Collection of its own pairs does not compile. With Keyed, the pair type
// would grow without end, an instantiation cycle.
func Pairs[K comparable, V any](c KeyedCollection[K, V]) Collection[Pair[K, V]] {
	return Collection[Pair[K, V]]{seq: func(yield func(Pair[K, V]) bool) {
		for k, v := range c.Seq() {
			if !yield(Pair[K, V]{k, v}) {
				return
			}
		}
	}}
}

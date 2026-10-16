package collect

// GroupBy groups the values by the key f returns for each. The groups are
// in the order their keys first come. It reads every value before it
// yields the first group. C is always Collection[T], and Go infers it, as
// with [Collection.Chunk].
func (c Collection[T]) GroupBy[K comparable, C interface{ Collection[T] }](f func(T) K) KeyedCollection[K, C] {
	return KeyedCollection[K, C]{seq: func(yield func(K, C) bool) {
		keys, groups := groupsOf(func(add func(K, T)) {
			for v := range c.Seq() {
				add(f(v), v)
			}
		})
		for i, k := range keys {
			if !yield(k, C(FromSlice(groups[i]))) {
				return
			}
		}
	}}
}

// GroupBy groups the pairs by the key f returns for each. Each group keeps
// the pairs' own keys, as Laravel's groupBy with preserveKeys does. It
// reads every pair before it yields the first group. C is always
// KeyedCollection[K, V], and Go infers it.
func (c KeyedCollection[K, V]) GroupBy[G comparable, C interface{ KeyedCollection[K, V] }](f func(K, V) G) KeyedCollection[G, C] {
	return KeyedCollection[G, C]{seq: func(yield func(G, C) bool) {
		keys, groups := groupsOf(func(add func(G, Entry[K, V])) {
			for k, v := range c.Seq() {
				add(f(k, v), Entry[K, V]{k, v})
			}
		})
		for i, g := range keys {
			pairs := groups[i]
			members := KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
				for _, p := range pairs {
					if !yield(p.Key, p.Value) {
						return
					}
				}
			}}
			if !yield(g, C(members)) {
				return
			}
		}
	}}
}

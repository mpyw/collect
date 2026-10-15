package collect

// GroupBy groups the values by the key f returns for each. The groups are
// in the order their keys first come. It reads every value before it
// yields the first group. C is always Collection[T], and Go infers it, as
// with [Collection.Chunk].
func (c Collection[T]) GroupBy[K comparable, C interface{ Collection[T] }](f func(T) K) KeyedCollection[K, C] {
	return KeyedCollection[K, C]{seq: func(yield func(K, C) bool) {
		groups := map[K][]T{}
		var order []K
		for v := range c.Seq() {
			k := f(v)
			if _, ok := groups[k]; !ok {
				order = append(order, k)
			}
			groups[k] = append(groups[k], v)
		}
		for _, k := range order {
			if !yield(k, C(FromSlice(groups[k]))) {
				return
			}
		}
	}}
}

// GroupBy groups the pairs by the key f returns for each. Each group keeps
// the pairs' own keys, as Laravel's groupBy with preserveKeys does. C is
// always KeyedCollection[K, V], and Go infers it.
func (c KeyedCollection[K, V]) GroupBy[G comparable, C interface{ KeyedCollection[K, V] }](f func(K, V) G) KeyedCollection[G, C] {
	return KeyedCollection[G, C]{seq: func(yield func(G, C) bool) {
		type group struct {
			ks []K
			vs []V
		}
		groups := map[G]*group{}
		var order []G
		for k, v := range c.Seq() {
			g := f(k, v)
			if groups[g] == nil {
				groups[g] = &group{}
				order = append(order, g)
			}
			groups[g].ks, groups[g].vs = append(groups[g].ks, k), append(groups[g].vs, v)
		}
		for _, g := range order {
			ks, vs := groups[g].ks, groups[g].vs
			members := KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
				for i, k := range ks {
					if !yield(k, vs[i]) {
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

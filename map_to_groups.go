package collect

// MapToGroups groups the values f returns by the keys it returns with them.
// The groups are in the order their keys first come. It reads every value
// before it yields the first group. Laravel's mapToDictionary is the same
// in Go.
func (c Collection[T]) MapToGroups[K comparable, U any](f func(T) (K, U)) KeyedCollection[K, Collection[U]] {
	return KeyedCollection[K, Collection[U]]{seq: func(yield func(K, Collection[U]) bool) {
		groups := map[K][]U{}
		var order []K
		for v := range c.Seq() {
			k, u := f(v)
			if _, ok := groups[k]; !ok {
				order = append(order, k)
			}
			groups[k] = append(groups[k], u)
		}
		for _, k := range order {
			if !yield(k, FromSlice(groups[k])) {
				return
			}
		}
	}}
}

// MapToGroups groups the values f returns by the keys it returns with them,
// as [Collection.MapToGroups] does.
func (c KeyedCollection[K, V]) MapToGroups[G comparable, U any](f func(K, V) (G, U)) KeyedCollection[G, Collection[U]] {
	return KeyedCollection[G, Collection[U]]{seq: func(yield func(G, Collection[U]) bool) {
		groups := map[G][]U{}
		var order []G
		for k, v := range c.Seq() {
			g, u := f(k, v)
			if _, ok := groups[g]; !ok {
				order = append(order, g)
			}
			groups[g] = append(groups[g], u)
		}
		for _, g := range order {
			if !yield(g, FromSlice(groups[g])) {
				return
			}
		}
	}}
}

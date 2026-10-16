package collect

// MapToGroups groups the values f returns by the keys it returns with them.
// The groups are in the order their keys first come. It reads every value
// before it yields the first group. Laravel's mapToDictionary is the same
// in Go.
func (c Collection[T]) MapToGroups[K comparable, U any](f func(T) (K, U)) KeyedCollection[K, Collection[U]] {
	return KeyedCollection[K, Collection[U]]{seq: func(yield func(K, Collection[U]) bool) {
		keys, groups := groupsOf(func(add func(K, U)) {
			for v := range c.Seq() {
				add(f(v))
			}
		})
		for i, k := range keys {
			if !yield(k, FromSlice(groups[i])) {
				return
			}
		}
	}}
}

// MapToGroups groups the values f returns by the keys it returns with them,
// as [Collection.MapToGroups] does.
func (c KeyedCollection[K, V]) MapToGroups[G comparable, U any](f func(K, V) (G, U)) KeyedCollection[G, Collection[U]] {
	return KeyedCollection[G, Collection[U]]{seq: func(yield func(G, Collection[U]) bool) {
		keys, groups := groupsOf(func(add func(G, U)) {
			for k, v := range c.Seq() {
				add(f(k, v))
			}
		})
		for i, g := range keys {
			if !yield(g, FromSlice(groups[i])) {
				return
			}
		}
	}}
}

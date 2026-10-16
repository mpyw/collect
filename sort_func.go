package collect

import "slices"

// SortFunc returns the values sorted by cmp, which returns a negative
// number, zero or a positive number, as [cmp.Compare] does. Equal values
// keep their order. It reads every value before the first.
func (c Collection[T]) SortFunc(cmp func(a, b T) int) Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		all := c.All()
		slices.SortStableFunc(all, cmp)
		for _, v := range all {
			if !yield(v) {
				return
			}
		}
	}}
}

// SortFunc returns the pairs sorted by cmp, with their keys, as
// [Collection.SortFunc] sorts values.
func (c KeyedCollection[K, V]) SortFunc(cmp func(ak K, av V, bk K, bv V) int) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		var pairs []Entry[K, V]
		for k, v := range c.Seq() {
			pairs = append(pairs, Entry[K, V]{k, v})
		}
		slices.SortStableFunc(pairs, func(a, b Entry[K, V]) int { return cmp(a.Key, a.Value, b.Key, b.Value) })
		for _, p := range pairs {
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}}
}

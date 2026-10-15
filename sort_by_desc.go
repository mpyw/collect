package collect

import "cmp"

// SortByDesc returns the values sorted by the key f returns for each,
// greatest first. Equal keys keep their order. Pass [Identity] for
// Laravel's sortDesc.
func (c Collection[T]) SortByDesc[K cmp.Ordered](f func(T) K) Collection[T] {
	return c.SortFunc(func(a, b T) int { return cmp.Compare(f(b), f(a)) })
}

// SortByDesc returns the pairs sorted by the key f returns for each,
// greatest first, with their own keys.
func (c KeyedCollection[K, V]) SortByDesc[O cmp.Ordered](f func(K, V) O) KeyedCollection[K, V] {
	return c.SortFunc(func(ak K, av V, bk K, bv V) int { return cmp.Compare(f(bk, bv), f(ak, av)) })
}

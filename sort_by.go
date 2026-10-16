package collect

import "cmp"

// SortBy returns the values sorted by the key f returns for each, least
// first. Equal keys keep their order. Pass [Identity] for Laravel's sort.
// It reads every value before the first.
func (c Collection[T]) SortBy[K cmp.Ordered](f func(T) K) Collection[T] {
	return c.SortFunc(func(a, b T) int { return cmp.Compare(f(a), f(b)) })
}

// SortBy returns the pairs sorted by the key f returns for each, least
// first, with their own keys. It reads every pair before the first.
func (c KeyedCollection[K, V]) SortBy[O cmp.Ordered](f func(K, V) O) KeyedCollection[K, V] {
	return c.SortFunc(func(ak K, av V, bk K, bv V) int { return cmp.Compare(f(ak, av), f(bk, bv)) })
}

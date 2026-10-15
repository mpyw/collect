package collect

import "cmp"

// SortKeysBy returns the pairs sorted by what f returns for each key, least
// first. Pass [Identity] for ordered keys.
func (c KeyedCollection[K, V]) SortKeysBy[O cmp.Ordered](f func(K) O) KeyedCollection[K, V] {
	return c.SortFunc(func(ak K, _ V, bk K, _ V) int { return cmp.Compare(f(ak), f(bk)) })
}

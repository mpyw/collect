package collect

// SortKeysFunc returns the pairs sorted by cmp of their keys. It is
// Laravel's sortKeysUsing.
func (c KeyedCollection[K, V]) SortKeysFunc(cmp func(a, b K) int) KeyedCollection[K, V] {
	return c.SortFunc(func(ak K, _ V, bk K, _ V) int { return cmp(ak, bk) })
}

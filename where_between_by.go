package collect

import "cmp"

// WhereBetweenBy returns the values whose key, as f returns it, is from low
// to high, both included.
func (c Collection[T]) WhereBetweenBy[K cmp.Ordered](f func(T) K, low, high K) Collection[T] {
	return c.Filter(func(v T) bool {
		k := f(v)
		return cmp.Compare(k, low) >= 0 && cmp.Compare(k, high) <= 0
	})
}

// WhereBetweenBy returns the pairs whose key, as f returns it, is from low to
// high, both included.
func (c KeyedCollection[K, V]) WhereBetweenBy[E cmp.Ordered](f func(K, V) E, low, high E) KeyedCollection[K, V] {
	return c.Filter(func(k K, v V) bool {
		e := f(k, v)
		return cmp.Compare(e, low) >= 0 && cmp.Compare(e, high) <= 0
	})
}

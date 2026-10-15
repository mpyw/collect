package collect

import "cmp"

// WhereNotBetweenBy returns the values whose key, as f returns it, is below
// low or above high.
func (c Collection[T]) WhereNotBetweenBy[K cmp.Ordered](f func(T) K, low, high K) Collection[T] {
	return c.Reject(func(v T) bool {
		k := f(v)
		return cmp.Compare(k, low) >= 0 && cmp.Compare(k, high) <= 0
	})
}

// WhereNotBetweenBy returns the pairs whose key, as f returns it, is below
// low or above high.
func (c KeyedCollection[K, V]) WhereNotBetweenBy[E cmp.Ordered](f func(K, V) E, low, high E) KeyedCollection[K, V] {
	return c.Reject(func(k K, v V) bool {
		e := f(k, v)
		return cmp.Compare(e, low) >= 0 && cmp.Compare(e, high) <= 0
	})
}

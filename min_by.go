package collect

import "cmp"

// MinBy returns the least of the keys f returns for the values. It is false
// for an empty collection. Pass [Identity] for a collection of keys.
func (c Collection[T]) MinBy[K cmp.Ordered](f func(T) K) (K, bool) {
	var (
		least K
		ok    bool
	)
	for v := range c.Seq() {
		if k := f(v); !ok || cmp.Less(k, least) {
			least, ok = k, true
		}
	}
	return least, ok
}

// MinBy returns the least of the keys f returns for the pairs. It is false
// for an empty collection.
func (c KeyedCollection[K, V]) MinBy[O cmp.Ordered](f func(K, V) O) (O, bool) {
	var (
		least O
		ok    bool
	)
	for k, v := range c.Seq() {
		if o := f(k, v); !ok || cmp.Less(o, least) {
			least, ok = o, true
		}
	}
	return least, ok
}

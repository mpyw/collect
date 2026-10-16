package collect

import "cmp"

// MaxBy returns the greatest of the keys f returns for the values, as
// Laravel's max returns it: the key, not the value. For the value itself,
// use SortByDesc(f).First(). It is false for an empty collection. Pass
// [Identity] for a collection of keys.
func (c Collection[T]) MaxBy[K cmp.Ordered](f func(T) K) (K, bool) {
	var (
		greatest K
		ok       bool
	)
	for v := range c.Seq() {
		if k := f(v); !ok || cmp.Less(greatest, k) {
			greatest, ok = k, true
		}
	}
	return greatest, ok
}

// MaxBy returns the greatest of the keys f returns for the pairs. It is
// false for an empty collection.
func (c KeyedCollection[K, V]) MaxBy[O cmp.Ordered](f func(K, V) O) (O, bool) {
	var (
		greatest O
		ok       bool
	)
	for k, v := range c.Seq() {
		if o := f(k, v); !ok || cmp.Less(greatest, o) {
			greatest, ok = o, true
		}
	}
	return greatest, ok
}

package collect

// ModeBy returns the keys that f returns most often for the values, in the
// order they first come. It is empty for an empty collection. Pass
// [Identity] for a collection of keys.
func (c Collection[T]) ModeBy[K comparable](f func(T) K) []K {
	counts := map[K]int{}
	var order []K
	most := 0
	for v := range c.Seq() {
		k := f(v)
		if counts[k] == 0 {
			order = append(order, k)
		}
		counts[k]++
		most = max(most, counts[k])
	}
	var modes []K
	for _, k := range order {
		if counts[k] == most {
			modes = append(modes, k)
		}
	}
	return modes
}

// ModeBy returns the keys that f returns most often for the pairs, in the
// order they first come. It is empty for an empty collection.
func (c KeyedCollection[K, V]) ModeBy[E comparable](f func(K, V) E) []E {
	var es []E
	for k, v := range c.Seq() {
		es = append(es, f(k, v))
	}
	return FromSlice(es).ModeBy(Identity)
}

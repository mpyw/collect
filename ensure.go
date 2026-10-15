package collect

// Ensure reads every value now, and returns them as U. The error wraps
// [ErrUnexpectedValue] and names the first value that is not a U. Write U:
// c.Ensure[*Dog](). It is not lazy: it has to read every value to know
// there is no error. [Collection.MustEnsure] is lazy, and panics instead.
func (c Collection[T]) Ensure[U any]() (Collection[U], error) {
	var all []U
	for v := range c.Seq() {
		u, err := asType[T, U](v)
		if err != nil {
			return Collection[U]{}, err
		}
		all = append(all, u)
	}
	return FromSlice(all), nil
}

// Ensure reads every pair now, and returns them with their values as U.
// The error wraps [ErrUnexpectedValue] and names the first value that is
// not a U. Write U: c.Ensure[*Dog]().
func (c KeyedCollection[K, V]) Ensure[U any]() (KeyedCollection[K, U], error) {
	var (
		ks []K
		us []U
	)
	for k, v := range c.Seq() {
		u, err := asType[V, U](v)
		if err != nil {
			return KeyedCollection[K, U]{}, err
		}
		ks, us = append(ks, k), append(us, u)
	}
	return keyedOf(ks, us), nil
}

package collect

// WhereInstanceOf returns the values that are of type U, as U. Write U:
// c.WhereInstanceOf[*Dog](). An interface U keeps the values that
// implement it.
func (c Collection[T]) WhereInstanceOf[U any]() Collection[U] {
	return Collection[U]{seq: func(yield func(U) bool) {
		for v := range c.Seq() {
			if u, ok := any(v).(U); ok && !yield(u) {
				return
			}
		}
	}}
}

// WhereInstanceOf returns the pairs whose value is of type U, with the
// value as U. Write U: c.WhereInstanceOf[*Dog]().
func (c KeyedCollection[K, V]) WhereInstanceOf[U any]() KeyedCollection[K, U] {
	return KeyedCollection[K, U]{seq: func(yield func(K, U) bool) {
		for k, v := range c.Seq() {
			if u, ok := any(v).(U); ok && !yield(k, u) {
				return
			}
		}
	}}
}

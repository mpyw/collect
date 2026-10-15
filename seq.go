package collect

import "iter"

// Seq returns the sequence of the collection, for a range loop or for the
// iter, slices and maps packages.
func (c Collection[T]) Seq() iter.Seq[T] {
	if c.seq == nil {
		return func(func(T) bool) {}
	}
	return c.seq
}

// Seq returns the sequence of the keys and values, for a range loop or for
// the iter, slices and maps packages.
func (c KeyedCollection[K, V]) Seq() iter.Seq2[K, V] {
	if c.seq == nil {
		return func(func(K, V) bool) {}
	}
	return c.seq
}

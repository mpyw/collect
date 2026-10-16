package collect

// Nth returns every step-th value, starting at offset. A negative offset
// counts from the end. It panics with [ErrInvalidArgument] if step is less than 1, as Laravel throws.
func (c Collection[T]) Nth(step, offset int) Collection[T] {
	if step < 1 {
		invalidArgument("Nth step must be at least 1, not %d", step)
	}
	from := c.Skip(offset)
	if offset < 0 {
		from = c.Take(offset)
	}
	return from.Keyed().
		Filter(func(i int, _ T) bool { return i%step == 0 }).
		Values()
}

// Nth returns every step-th pair, starting at offset, with their keys. A
// negative offset counts from the end. It panics with [ErrInvalidArgument] if step is less than 1.
func (c KeyedCollection[K, V]) Nth(step, offset int) KeyedCollection[K, V] {
	if step < 1 {
		invalidArgument("Nth step must be at least 1, not %d", step)
	}
	from := c.Skip(offset)
	if offset < 0 {
		from = c.Take(offset)
	}
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		i := 0
		for k, v := range from.Seq() {
			if i%step == 0 && !yield(k, v) {
				return
			}
			i++
		}
	}}
}

package collect

// Slice returns length values from offset, as PHP's array_slice does. A
// negative offset counts from the end. A negative length stops that many
// values before the end. With both non-negative it is Skip(offset).
// Take(length), and lazy. Otherwise it reads every value.
func (c Collection[T]) Slice(offset, length int) Collection[T] {
	if offset >= 0 && length >= 0 {
		return c.Skip(offset).Take(length)
	}
	return Collection[T]{seq: func(yield func(T) bool) {
		all := c.All()
		start, end := sliceBounds(len(all), offset, length)
		for _, v := range all[start:end] {
			if !yield(v) {
				return
			}
		}
	}}
}

// Slice returns length pairs from offset, as PHP's array_slice does, with
// their keys. Negative values count from the end, as in
// [Collection.Slice].
func (c KeyedCollection[K, V]) Slice(offset, length int) KeyedCollection[K, V] {
	if offset >= 0 && length >= 0 {
		return c.Skip(offset).Take(length)
	}
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		var (
			ks []K
			vs []V
		)
		for k, v := range c.Seq() {
			ks, vs = append(ks, k), append(vs, v)
		}
		start, end := sliceBounds(len(ks), offset, length)
		for i := start; i < end; i++ {
			if !yield(ks[i], vs[i]) {
				return
			}
		}
	}}
}

// sliceBounds returns the indexes of array_slice's result in a list of n.
func sliceBounds(n, offset, length int) (start, end int) {
	start = offset
	if start < 0 {
		start = max(n+start, 0)
	}
	start = min(start, n)
	end = start + length
	if length < 0 {
		end = n + length
	}
	return start, min(max(end, start), n)
}

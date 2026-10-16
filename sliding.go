package collect

import "slices"

// Sliding returns windows of size values, each step values after the one
// before. Only full windows are returned. It panics with [ErrInvalidArgument] if size or step is
// less than 1, as Laravel throws. C is always Collection[T], and Go infers
// it.
func (c Collection[T]) Sliding[C interface{ Collection[T] }](size, step int) Collection[C] {
	if size < 1 || step < 1 {
		invalidArgument("Sliding size and step must be at least 1, not %d and %d", size, step)
	}
	return Collection[C]{seq: func(yield func(C) bool) {
		var window []T
		skip := 0
		for v := range c.Seq() {
			if skip > 0 {
				skip--
				continue
			}
			if window = append(window, v); len(window) < size {
				continue
			}
			if !yield(C(Of(window...))) {
				return
			}
			if step >= size {
				window, skip = window[:0], step-size
				continue
			}
			window = append(window[:0], window[step:]...)
		}
	}}
}

// Sliding returns windows of size pairs, each step pairs after the one
// before, as [Collection.Sliding] does. Each pair keeps its own key. C is
// always KeyedCollection[K, V], and Go infers it.
func (c KeyedCollection[K, V]) Sliding[C interface{ KeyedCollection[K, V] }](size, step int) Collection[C] {
	if size < 1 || step < 1 {
		invalidArgument("Sliding size and step must be at least 1, not %d and %d", size, step)
	}
	return Collection[C]{seq: func(yield func(C) bool) {
		var (
			ks []K
			vs []V
		)
		skip := 0
		for k, v := range c.Seq() {
			if skip > 0 {
				skip--
				continue
			}
			if ks, vs = append(ks, k), append(vs, v); len(ks) < size {
				continue
			}
			if !yield(C(keyedOf(slices.Clone(ks), slices.Clone(vs)))) {
				return
			}
			if step >= size {
				ks, vs, skip = ks[:0], vs[:0], step-size
				continue
			}
			ks, vs = append(ks[:0], ks[step:]...), append(vs[:0], vs[step:]...)
		}
	}}
}

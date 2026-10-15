package collect

import "fmt"

// Combine reads both now, and returns the keys, each with the value of
// values at the same place. The error wraps [ErrLengthMismatch] if there
// are not as many values as keys. It is not lazy: it has to read both to
// know there is no error. [MustCombine] is lazy, and panics instead.
func Combine[K comparable, V any](keys Collection[K], values Collection[V]) (KeyedCollection[K, V], error) {
	ks, vs := keys.All(), values.All()
	if len(ks) != len(vs) {
		return KeyedCollection[K, V]{}, fmt.Errorf("%w: %d keys and %d values", ErrLengthMismatch, len(ks), len(vs))
	}
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		for i, k := range ks {
			if !yield(k, vs[i]) {
				return
			}
		}
	}}, nil
}

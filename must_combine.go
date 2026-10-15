package collect

import (
	"fmt"
	"iter"
)

// MustCombine returns the keys, each with the value of values at the same
// place, lazily. Reading it panics with an error that wraps
// [ErrLengthMismatch] once one runs out before the other, as Laravel's
// combine throws.
func MustCombine[K comparable, V any](keys Collection[K], values Collection[V]) KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		next, stop := iter.Pull(values.Seq())
		defer stop()
		for k := range keys.Seq() {
			v, ok := next()
			if !ok {
				panic(fmt.Errorf("%w: more keys than values", ErrLengthMismatch))
			}
			if !yield(k, v) {
				return
			}
		}
		if _, ok := next(); ok {
			panic(fmt.Errorf("%w: more values than keys", ErrLengthMismatch))
		}
	}}
}

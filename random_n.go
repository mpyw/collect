package collect

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

// RandomN reads every value now, and returns n of them chosen at random,
// in the order they come. The error wraps [ErrNotEnoughItems] if there are
// fewer than n. A negative n is 0, as in Laravel.
func (c Collection[T]) RandomN(n int) (Collection[T], error) {
	n = max(n, 0)
	all := c.All()
	if len(all) < n {
		return Collection[T]{}, fmt.Errorf("%w: asked for %d, but there are only %d", ErrNotEnoughItems, n, len(all))
	}
	picked := rand.Perm(len(all))[:n]
	slices.Sort(picked)
	chosen := make([]T, n)
	for i, p := range picked {
		chosen[i] = all[p]
	}
	return FromSlice(chosen), nil
}

// RandomN reads every pair now, and returns n of them chosen at random, in
// the order they come, with their keys. The error wraps [ErrNotEnoughItems]
// if there are fewer than n. A negative n is 0, as in Laravel.
func (c KeyedCollection[K, V]) RandomN(n int) (KeyedCollection[K, V], error) {
	n = max(n, 0)
	ks, vs := pairsOf(c)
	if len(ks) < n {
		return KeyedCollection[K, V]{}, fmt.Errorf("%w: asked for %d, but there are only %d", ErrNotEnoughItems, n, len(ks))
	}
	picked := rand.Perm(len(ks))[:n]
	slices.Sort(picked)
	chosenK, chosenV := make([]K, n), make([]V, n)
	for i, p := range picked {
		chosenK[i], chosenV[i] = ks[p], vs[p]
	}
	return keyedOf(chosenK, chosenV), nil
}

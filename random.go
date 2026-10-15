package collect

import "math/rand/v2"

// Random returns a value chosen at random. It is false for an empty
// collection. It reads every value, and keeps only one.
func (c Collection[T]) Random() (T, bool) {
	var (
		chosen T
		n      int
	)
	for v := range c.Seq() {
		// Reservoir sampling: the nth value replaces the chosen one with
		// probability 1/n, so each is chosen with probability 1/count.
		if n++; rand.IntN(n) == 0 {
			chosen = v
		}
	}
	return chosen, n > 0
}

// Random returns a pair chosen at random. It is false for an empty
// collection. It reads every pair, and keeps only one.
func (c KeyedCollection[K, V]) Random() (K, V, bool) {
	var (
		chosenK K
		chosenV V
		n       int
	)
	for k, v := range c.Seq() {
		if n++; rand.IntN(n) == 0 {
			chosenK, chosenV = k, v
		}
	}
	return chosenK, chosenV, n > 0
}

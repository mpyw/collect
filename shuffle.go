package collect

import "math/rand/v2"

// Shuffle returns the values in a random order, new on each read. It reads
// every value before the first.
func (c Collection[T]) Shuffle() Collection[T] {
	return Collection[T]{seq: func(yield func(T) bool) {
		all := c.All()
		rand.Shuffle(len(all), func(i, j int) { all[i], all[j] = all[j], all[i] })
		for _, v := range all {
			if !yield(v) {
				return
			}
		}
	}}
}

// Shuffle returns the pairs in a random order, new on each read, with their
// keys. It reads every pair before the first.
func (c KeyedCollection[K, V]) Shuffle() KeyedCollection[K, V] {
	return KeyedCollection[K, V]{seq: func(yield func(K, V) bool) {
		ks, vs := pairsOf(c)
		rand.Shuffle(len(ks), func(i, j int) {
			ks[i], ks[j] = ks[j], ks[i]
			vs[i], vs[j] = vs[j], vs[i]
		})
		for k, v := range keyedOf(ks, vs).Seq() {
			if !yield(k, v) {
				return
			}
		}
	}}
}

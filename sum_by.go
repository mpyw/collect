package collect

// SumBy returns the sum of the numbers f returns for the values. Pass
// [Identity] for a collection of numbers.
func (c Collection[T]) SumBy[N Number](f func(T) N) N {
	var sum N
	for v := range c.Seq() {
		sum += f(v)
	}
	return sum
}

// SumBy returns the sum of the numbers f returns for the pairs.
func (c KeyedCollection[K, V]) SumBy[N Number](f func(K, V) N) N {
	var sum N
	for k, v := range c.Seq() {
		sum += f(k, v)
	}
	return sum
}

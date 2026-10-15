package collect

// AvgBy returns the mean of the numbers f returns for the values. It is false
// for an empty collection. Pass [Identity] for a collection of numbers.
func (c Collection[T]) AvgBy[N Number](f func(T) N) (float64, bool) {
	var (
		sum float64
		n   int
	)
	for v := range c.Seq() {
		sum += float64(f(v))
		n++
	}
	if n == 0 {
		return 0, false
	}
	return sum / float64(n), true
}

// AvgBy returns the mean of the numbers f returns for the pairs. It is false
// for an empty collection.
func (c KeyedCollection[K, V]) AvgBy[N Number](f func(K, V) N) (float64, bool) {
	var (
		sum float64
		n   int
	)
	for k, v := range c.Seq() {
		sum += float64(f(k, v))
		n++
	}
	if n == 0 {
		return 0, false
	}
	return sum / float64(n), true
}

package collect

// Reduce returns f applied to the result so far and each value in turn,
// starting from initial. The arguments are in Laravel's order.
func (c Collection[T]) Reduce[A any](f func(A, T) A, initial A) A {
	result := initial
	for v := range c.Seq() {
		result = f(result, v)
	}
	return result
}

// Reduce returns f applied to the result so far and each pair in turn,
// starting from initial. It is Laravel's reduceWithKeys.
func (c KeyedCollection[K, V]) Reduce[A any](f func(A, K, V) A, initial A) A {
	result := initial
	for k, v := range c.Seq() {
		result = f(result, k, v)
	}
	return result
}

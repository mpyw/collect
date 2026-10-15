package collect

// Times returns what f returns for 1 to n, in order, as Laravel's times
// does. It is empty if n is less than 1. Use [Range] for the numbers
// themselves.
func Times[T any](n int, f func(int) T) Collection[T] {
	return Range(1, n).Take(max(n, 0)).Map(f)
}

package collect

// JoinBy is [Collection.ImplodeBy], under the name of [strings.Join]. For
// Laravel's join with a final glue, see [Collection.JoinFinalBy].
func (c Collection[T]) JoinBy[S ~string](f func(T) S, sep string) string {
	return c.ImplodeBy(f, sep)
}

// JoinBy is [KeyedCollection.ImplodeBy], under the name of [strings.Join].
func (c KeyedCollection[K, V]) JoinBy[S ~string](f func(K, V) S, sep string) string {
	return c.ImplodeBy(f, sep)
}

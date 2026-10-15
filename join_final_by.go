package collect

// JoinFinalBy joins the strings f returns for the values, with glue between
// them and finalGlue before the last, as in "a, b and c". It is Laravel's
// join with a final glue. An empty finalGlue is glue. Pass [Identity] for a
// collection of strings.
func (c Collection[T]) JoinFinalBy[S ~string](f func(T) S, glue, finalGlue string) string {
	if finalGlue == "" {
		return c.ImplodeBy(f, glue)
	}
	values := c.Map(func(v T) string { return string(f(v)) }).All()
	switch n := len(values); n {
	case 0:
		return ""
	case 1:
		return values[0]
	default:
		return FromSlice(values[:n-1]).ImplodeBy(Identity, glue) + finalGlue + values[n-1]
	}
}

// JoinFinalBy joins the strings f returns for the pairs, with glue between
// them and finalGlue before the last, as [Collection.JoinFinalBy] does.
func (c KeyedCollection[K, V]) JoinFinalBy[S ~string](f func(K, V) S, glue, finalGlue string) string {
	var parts []string
	for k, v := range c.Seq() {
		parts = append(parts, string(f(k, v)))
	}
	return FromSlice(parts).JoinFinalBy(Identity, glue, finalGlue)
}

package collect

import "strings"

// ImplodeBy joins the strings f returns for the values, with sep between
// them. Pass [Identity] for a collection of strings.
func (c Collection[T]) ImplodeBy[S ~string](f func(T) S, sep string) string {
	return strings.Join(c.Map(func(v T) string { return string(f(v)) }).All(), sep)
}

// ImplodeBy joins the strings f returns for the pairs, with sep between them.
func (c KeyedCollection[K, V]) ImplodeBy[S ~string](f func(K, V) S, sep string) string {
	var parts []string
	for k, v := range c.Seq() {
		parts = append(parts, string(f(k, v)))
	}
	return strings.Join(parts, sep)
}

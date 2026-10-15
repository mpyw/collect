package collect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's diffUsing compares with a function.
func ExampleCollection_DiffFunc() {
	fmt.Println(collect.Of("A", "b", "C").DiffFunc(collect.Of("a", "c"), strings.EqualFold).All())
	// Output: [b]
}

func TestDiffByFuncStopsEarly(t *testing.T) {
	never := func(a, b int) bool { return false }
	assertStopsEarly(t, endless().DiffFunc(collect.Of(2), never).Seq())
}

func ExampleKeyedCollection_DiffFunc() {
	c := collect.Of("A", "b").Keyed().DiffFunc(collect.Of("a").Keyed(), strings.EqualFold)

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [1] [b]
}

func TestKeyedDiffByFuncStopsEarly(t *testing.T) {
	never := func(a, b int) bool { return false }
	assertStopsEarly2(t, endless().Keyed().DiffFunc(collect.Of(2).Keyed(), never).Seq())
}

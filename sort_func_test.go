package collect_test

import (
	"cmp"
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's sort with a callback. Equal values keep their order.
func ExampleCollection_SortFunc() {
	byLength := func(a, b string) int { return cmp.Compare(len(a), len(b)) }

	fmt.Println(collect.Of("ccc", "a", "bb", "d").SortFunc(byLength).All())
	// Output: [a d bb ccc]
}

func ExampleKeyedCollection_SortFunc() {
	byValue := func(_ int, a string, _ int, b string) int { return cmp.Compare(a, b) }
	c := collect.Of("b", "a").Keyed().SortFunc(byValue)

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [1 0] [a b]
}

func TestSortFuncStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Range(1, 9).SortFunc(cmp.Compare[int]).Seq())
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().SortFunc(func(a, _, b, _ int) int { return cmp.Compare(a, b) }).Seq())
}

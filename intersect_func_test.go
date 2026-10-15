package collect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's intersectUsing compares with a function.
func ExampleCollection_IntersectFunc() {
	fmt.Println(collect.Of("A", "b", "C").IntersectFunc(collect.Of("a", "c"), strings.EqualFold).All())
	// Output: [A C]
}

func TestIntersectByFuncStopsEarly(t *testing.T) {
	always := func(a, b int) bool { return true }
	assertStopsEarly(t, endless().IntersectFunc(collect.Of(2), always).Seq())
}

func ExampleKeyedCollection_IntersectFunc() {
	c := collect.Of("A", "b").Keyed().IntersectFunc(collect.Of("a").Keyed(), strings.EqualFold)

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [0] [A]
}

func TestKeyedIntersectByFuncStopsEarly(t *testing.T) {
	always := func(a, b int) bool { return true }
	assertStopsEarly2(t, endless().Keyed().IntersectFunc(collect.Of(2).Keyed(), always).Seq())
}

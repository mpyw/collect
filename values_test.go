package collect_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleKeyedCollection_Values() {
	c := collect.FromSeq2(slices.All([]string{"a", "b"})).
		Filter(func(i int, _ string) bool { return i > 0 })

	fmt.Println(c.Values().All())
	// Output: [b]
}

func TestValuesStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Keyed().Values().Seq())
}

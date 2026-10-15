package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Filter() {
	big := collect.Of(1, 2, 3, 4).
		Filter(func(n int) bool { return n > 2 }).
		All()

	fmt.Println(big)
	// Output: [3 4]
}

// Filter keeps the keys, and Values drops them: Laravel's filter and
// values, made explicit.
func ExampleKeyedCollection_Filter() {
	odd := collect.Of(10, 11, 12, 13).Keyed().
		Filter(func(_, n int) bool { return n%2 == 1 })

	fmt.Println(odd.Keys().All(), odd.Values().All())
	// Output: [1 3] [11 13]
}

func TestFilterStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Filter(func(int) bool { return true }).Seq())
	assertStopsEarly2(t, endless().Keyed().Filter(func(int, int) bool { return true }).Seq())
}

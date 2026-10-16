package collect_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of pad.
func ExampleCollection_Pad() {
	c := collect.Of("A", "B", "C")

	fmt.Println(c.Pad(5, "0").All())
	fmt.Println(c.Pad(-5, "0").All())
	// Output:
	// [A B C 0 0]
	// [0 0 A B C]
}

func TestPadShorter(t *testing.T) {
	c := collect.Of("A", "B", "C")
	assertEqual(t, c.Pad(2, "0").All(), []string{"A", "B", "C"})
	assertEqual(t, c.Pad(-2, "0").All(), []string{"A", "B", "C"})
}

func TestPadStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Of(1).Pad(10, 0).Seq())
	assertStopsEarly(t, collect.Of(1, 2, 3, 4).Pad(2, 0).Seq())
	assertStopsEarly(t, collect.Of(1).Pad(-10, 0).Seq())
	assertStopsEarly(t, collect.Of(1, 2, 3, 4).Pad(-6, 0).Seq())
}

// math.MinInt, whose negation overflows, pads as far as math.MaxInt.
func TestPadMinInt(t *testing.T) {
	assertEqual(t, collect.Of(1).Pad(math.MinInt, 0).Take(3).All(), []int{0, 0, 0})
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of times.
func ExampleTimes() {
	fmt.Println(collect.Times(5, func(n int) int { return n * 9 }).All())
	// Output: [9 18 27 36 45]
}

func TestTimesNone(t *testing.T) {
	assertEqual(t, collect.Times(0, collect.Identity[int]).All(), nil)
	assertEqual(t, collect.Times(-2, collect.Identity[int]).All(), nil)
}

func TestTimesStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Times(9, collect.Identity[int]).Seq())
}

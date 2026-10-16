package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Zip3() {
	names := collect.Of("Chair", "Desk")
	prices := collect.Of(100, 200)
	stock := collect.Of(true, false)

	for t := range names.Zip3(prices, stock).Seq() {
		fmt.Println(t.A, t.B, t.C)
	}
	// Output:
	// Chair 100 true
	// Desk 200 false
}

// Zip3 goes on to the end of the longest, with zero values for the rest.
func TestZip3Uneven(t *testing.T) {
	got := collect.Of(1).Zip3(collect.Of("a", "b"), collect.Of[bool]()).All()
	assertEqual(t, got, []collect.Tuple3[int, string, bool]{{1, "a", false}, {0, "b", false}})
}

func TestZip3StopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Zip3(endless(), endless()).Seq())
}

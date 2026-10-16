package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of zip.
func ExampleCollection_Zip2() {
	for p := range collect.Of("Chair", "Desk").Zip2(collect.Of(100, 200)).Seq() {
		fmt.Println(p.A, p.B)
	}
	// Output:
	// Chair 100
	// Desk 200
}

// Zip2 goes on to the end of the longer one, as Laravel does.
func TestZip2Uneven(t *testing.T) {
	pairs := collect.Of("Chair", "Desk").Zip2(collect.Of(100, 200, 300)).All()
	assertEqual(t, pairs, []collect.Tuple2[string, int]{{"Chair", 100}, {"Desk", 200}, {"", 300}})
	assertEqual(t, collect.Of(1, 2).Zip2(collect.Of[string]()).All(), []collect.Tuple2[int, string]{{1, ""}, {2, ""}})
}

func TestZip2StopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Zip2(endless()).Seq())
}

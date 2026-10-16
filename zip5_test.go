package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Zip5() {
	t := collect.Of(1).Zip5(collect.Of("a"), collect.Of(true), collect.Of(1.5), collect.Of('x')).MustFirst()

	fmt.Println(t.A, t.B, t.C, t.D, string(t.E))
	// Output: 1 a true 1.5 x
}

func TestZip5Uneven(t *testing.T) {
	got := collect.Of(1).Zip5(collect.Of[string](), collect.Of[bool](), collect.Of[float64](), collect.Of('x', 'y')).All()
	want := []collect.Tuple5[int, string, bool, float64, rune]{{1, "", false, 0, 'x'}, {0, "", false, 0, 'y'}}
	assertEqual(t, got, want)
}

func TestZip5StopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Zip5(endless(), endless(), endless(), endless()).Seq())
}

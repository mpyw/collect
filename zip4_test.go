package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Zip4() {
	t := collect.Of(1).Zip4(collect.Of("a"), collect.Of(true), collect.Of(1.5)).MustFirst()

	fmt.Println(t.A, t.B, t.C, t.D)
	// Output: 1 a true 1.5
}

func TestZip4Uneven(t *testing.T) {
	got := collect.Of(1, 2).Zip4(collect.Of("a"), collect.Of[bool](), collect.Of(1.5, 2.5, 3.5)).All()
	want := []collect.Tuple4[int, string, bool, float64]{{1, "a", false, 1.5}, {2, "", false, 2.5}, {0, "", false, 3.5}}
	assertEqual(t, got, want)
}

func TestZip4StopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Zip4(endless(), endless(), endless()).Seq())
}

package collect_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleRangeStep() {
	fmt.Println(collect.RangeStep(0, 10, 3).All())
	fmt.Println(collect.RangeStep(0, 1, 0.25).All())
	// Output:
	// [0 3 6 9]
	// [0 0.25 0.5 0.75 1]
}

func TestRangeStep(t *testing.T) {
	assertEqual(t, collect.RangeStep(10, 0, -3).All(), []int{10, 7, 4, 1})
	assertEqual(t, collect.RangeStep(10, 0, 3).All(), []int{10, 7, 4, 1})
	assertEqual(t, collect.RangeStep[uint8](0, 255, 100).All(), []uint8{0, 100, 200})
	assertEqual(t, collect.RangeStep[int8](-128, 127, 100).All(), []int8{-128, -28, 72})
	assertPanicsWith(t, collect.ErrInvalidArgument, func() { collect.RangeStep(0, 1, 0) })
}

// A float range computes each number from from, as PHP's range does, so
// that it ends on its last number exactly.
func TestRangeStepFloat(t *testing.T) {
	assertEqual(t, collect.RangeStep(0, 1, 0.1).All(), []float64{0, 0.1, 0.2, 0.30000000000000004, 0.4, 0.5, 0.6000000000000001, 0.7000000000000001, 0.8, 0.9, 1})
	assertEqual(t, collect.RangeStep(1, 0, 0.1).All(), []float64{1, 0.9, 0.8, 0.7, 0.6, 0.5, 0.3999999999999999, 0.29999999999999993, 0.19999999999999996, 0.09999999999999998, 0})
	assertEqual(t, collect.RangeStep(0.9, 1, 0.1).All(), []float64{0.9, 1})
	assertEqual(t, collect.RangeStep(1, 0.9, 0.1).All(), []float64{1, 0.9})
	// 0.1 + 3*0.2 is 0.7000000000000001, past 0.7, in PHP too.
	assertEqual(t, collect.RangeStep(0.1, 0.7, 0.2).All(), []float64{0.1, 0.30000000000000004, 0.5})
	assertEqual(t, collect.RangeStep(0, 1, 0.3).All(), []float64{0, 0.3, 0.6, 0.8999999999999999})
	assertEqual(t, collect.RangeStep(1, 0, 0.3).All(), []float64{1, 0.7, 0.4, 0.10000000000000009})
	assertEqual(t, collect.RangeStep(0.9, 1, 0.05).All(), []float64{0.9, 0.9500000000000001, 1})
	if n := collect.RangeStep(0.0, 1, 0.1).Count(); n != 11 {
		t.Errorf("RangeStep(0, 1, 0.1) has %d numbers, want 11", n)
	}
}

// A step that changes nothing ends the range, and a NaN gives none.
func TestRangeStepEnds(t *testing.T) {
	assertEqual(t, collect.RangeStep(1.0, 10, 1e-300).All(), []float64{1})
	assertEqual(t, collect.RangeStep(1e300, 2e300, 1).All(), []float64{1e300})
	nan := math.NaN()
	for _, c := range []collect.Collection[float64]{
		collect.RangeStep(0, 1, nan),
		collect.Range(nan, 1),
		collect.Range(0, nan),
	} {
		assertEqual(t, c.All(), nil)
	}
}

func TestRangeStepStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.RangeStep(0, 100, 3).Seq())
	assertStopsEarly(t, collect.RangeStep(0, 10, 0.5).Seq())
}

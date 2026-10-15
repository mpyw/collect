package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of range, and counting down.
func ExampleRange() {
	fmt.Println(collect.Range(3, 6).All())
	fmt.Println(collect.Range(6, 3).All())
	// Output:
	// [3 4 5 6]
	// [6 5 4 3]
}

func ExampleRangeStep() {
	fmt.Println(collect.RangeStep(0, 10, 3).All())
	fmt.Println(collect.RangeStep(0, 1, 0.25).All())
	// Output:
	// [0 3 6 9]
	// [0 0.25 0.5 0.75 1]
}

// A step never wraps around the end of an unsigned type.
func TestRange(t *testing.T) {
	assertEqual(t, collect.Range(2, 2).All(), []int{2})
	assertEqual(t, collect.Range[uint8](253, 255).All(), []uint8{253, 254, 255})
	assertEqual(t, collect.Range[uint8](2, 0).All(), []uint8{2, 1, 0})
}

func TestRangeStep(t *testing.T) {
	assertEqual(t, collect.RangeStep(10, 0, -3).All(), []int{10, 7, 4, 1})
	assertEqual(t, collect.RangeStep(10, 0, 3).All(), []int{10, 7, 4, 1})
	assertEqual(t, collect.RangeStep[uint8](0, 255, 100).All(), []uint8{0, 100, 200})
	assertPanics(t, func() { collect.RangeStep(0, 1, 0) })
}

func TestRangeStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Range(1, 9).Seq())
	assertStopsEarly(t, collect.Range(9, 1).Seq())
}

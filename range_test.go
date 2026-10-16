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

// A range never wraps around the end of its type.
func TestRange(t *testing.T) {
	assertEqual(t, collect.Range(2, 2).All(), []int{2})
	assertEqual(t, collect.Range[uint8](253, 255).All(), []uint8{253, 254, 255})
	assertEqual(t, collect.Range[uint8](2, 0).All(), []uint8{2, 1, 0})
	if n := collect.Range[int8](-128, 127).Count(); n != 256 {
		t.Errorf("Range(-128, 127) of int8 has %d numbers, want 256", n)
	}
	if n := collect.Range[int16](-20000, 20000).Count(); n != 40001 {
		t.Errorf("Range(-20000, 20000) of int16 has %d numbers, want 40001", n)
	}
	if n := collect.Range[int8](100, -100).Count(); n != 201 {
		t.Errorf("Range(100, -100) of int8 has %d numbers, want 201", n)
	}
}

func TestRangeStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Range(1, 9).Seq())
	assertStopsEarly(t, collect.Range(9, 1).Seq())
}

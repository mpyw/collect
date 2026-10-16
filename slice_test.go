package collect_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of slice, and PHP's array_slice with negative values.
func ExampleCollection_Slice() {
	c := collect.Range(1, 10)

	fmt.Println(c.Slice(4, 2).All())
	fmt.Println(c.Slice(-3, 2).All())
	fmt.Println(c.Slice(2, -3).All())
	// Output:
	// [5 6]
	// [8 9]
	// [3 4 5 6 7]
}

func ExampleKeyedCollection_Slice() {
	fmt.Println(collect.Range(1, 10).Keyed().Slice(-2, 1).Keys().All())
	// Output: [8]
}

// Laravel's examples of slice, and PHP's array_slice with negative values.
func TestSlice(t *testing.T) {
	c := collect.Range(1, 10)
	for _, tt := range []struct {
		offset, length int
		want           []int
	}{
		{4, 2, []int{5, 6}},
		{4, 100, []int{5, 6, 7, 8, 9, 10}},
		{-3, 2, []int{8, 9}},
		{2, -3, []int{3, 4, 5, 6, 7}},
		{-2, -1, []int{9}},
		{-20, 2, []int{1, 2}},
		{20, -1, nil},
		{8, -5, nil},
	} {
		assertEqual(t, c.Slice(tt.offset, tt.length).All(), tt.want)
		assertEqual(t, c.Keyed().Slice(tt.offset, tt.length).Values().All(), tt.want)
	}
	assertEqual(t, c.Keyed().Slice(-2, 1).Keys().All(), []int{8})
}

func TestSliceStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Slice(1, 100).Seq())
	assertStopsEarly(t, collect.Range(1, 9).Slice(-8, 5).Seq())
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().Slice(-8, 5).Seq())
}

func TestSliceMaxInt(t *testing.T) {
	assertEqual(t, collect.Range(1, 5).Slice(-3, math.MaxInt).All(), []int{3, 4, 5})
}

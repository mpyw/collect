package collect_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of sliding.
func ExampleCollection_Sliding() {
	all := collect.Collection[int].All
	fmt.Println(collect.Range(1, 5).Sliding(2, 1).Map(all).All())
	fmt.Println(collect.Range(1, 5).Sliding(3, 2).Map(all).All())
	// Output:
	// [[1 2] [2 3] [3 4] [4 5]]
	// [[1 2 3] [3 4 5]]
}

// A step past the size skips values between windows.
func TestSliding(t *testing.T) {
	assertChunks(t, collect.Range(1, 7).Sliding(2, 3), [][]int{{1, 2}, {4, 5}})
	assertChunks(t, collect.Range(1, 4).Sliding(2, 2), [][]int{{1, 2}, {3, 4}})
	assertChunks(t, collect.Range(1, 2).Sliding(3, 1), nil)
	assertPanics(t, func() { collect.Of(1).Sliding(0, 1) })
	assertPanics(t, func() { collect.Of(1).Sliding(1, 0) })
}

// A window stays the same after the next one is made.
func TestSlidingKeepsEachWindow(t *testing.T) {
	all := collect.Range(1, 4).Sliding(2, 1).All()
	assertEqual(t, all[0].All(), []int{1, 2})
	assertEqual(t, all[2].All(), []int{3, 4})
}

func TestSlidingStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Sliding(2, 1).Seq())
}

func ExampleKeyedCollection_Sliding() {
	for window := range collect.Of("a", "b", "c").Keyed().Sliding(2, 1).Seq() {
		fmt.Println(window.Keys().All())
	}
	// Output:
	// [0 1]
	// [1 2]
}

func TestKeyedSliding(t *testing.T) {
	keys := func(c collect.KeyedCollection[int, int]) []int { return c.Keys().All() }
	got := collect.Range(1, 7).Keyed().Sliding(2, 3).Map(keys).All()
	if want := [][]int{{0, 1}, {3, 4}}; !slices.EqualFunc(got, want, slices.Equal) {
		t.Errorf("Sliding(2, 3) = %v, want %v", got, want)
	}
	assertPanics(t, func() { collect.Of(1).Keyed().Sliding(0, 1) })
	assertStopsEarly(t, endless().Keyed().Sliding(2, 1).Seq())
}

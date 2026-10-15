package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of split: the groups are as even as can be.
func ExampleCollection_Split() {
	fmt.Println(collect.Range(1, 5).Split(3).Map(collect.Collection[int].All).All())
	// Output: [[1 2] [3 4] [5]]
}

func TestSplit(t *testing.T) {
	assertChunks(t, collect.Range(1, 5).Split(3), [][]int{{1, 2}, {3, 4}, {5}})
	assertChunks(t, collect.Range(1, 7).Split(3), [][]int{{1, 2, 3}, {4, 5}, {6, 7}})
	assertChunks(t, collect.Range(1, 2).Split(3), [][]int{{1}, {2}})
	assertChunks(t, collect.Of[int]().Split(3), nil)
	assertPanics(t, func() { collect.Of(1).Split(0) })
}

func TestSplitStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Range(1, 9).Split(5).Seq())
}

func ExampleKeyedCollection_Split() {
	for group := range collect.Range(1, 5).Keyed().Split(3).Seq() {
		fmt.Println(group.Keys().All())
	}
	// Output:
	// [0 1]
	// [2 3]
	// [4]
}

func TestKeyedSplit(t *testing.T) {
	if n := collect.Of[int]().Keyed().Split(3).Count(); n != 0 {
		t.Errorf("Split() of nothing has %d groups", n)
	}
	assertPanics(t, func() { collect.Of(1).Keyed().Split(0) })
	assertStopsEarly(t, collect.Range(1, 9).Keyed().Split(5).Seq())
}

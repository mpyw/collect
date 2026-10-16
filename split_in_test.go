package collect_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of splitIn: each group is full but the last.
func ExampleCollection_SplitIn() {
	fmt.Println(collect.Range(1, 10).SplitIn(3).Map(collect.Collection[int].All).All())
	// Output: [[1 2 3 4] [5 6 7 8] [9 10]]
}

// Laravel's example of splitIn, and how it differs from split.
func TestSplitIn(t *testing.T) {
	assertChunks(t, collect.Range(1, 10).SplitIn(3), [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10}})
	assertChunks(t, collect.Range(1, 7).SplitIn(3), [][]int{{1, 2, 3}, {4, 5, 6}, {7}})
	assertChunks(t, collect.Of[int]().SplitIn(3), nil)
	assertPanicsWith(t, collect.ErrInvalidArgument, func() { collect.Of(1).SplitIn(0) })
}

func TestSplitInStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Range(1, 9).SplitIn(9).Seq())
}

func ExampleKeyedCollection_SplitIn() {
	for group := range collect.Range(1, 5).Keyed().SplitIn(2).Seq() {
		fmt.Println(group.Keys().All())
	}
	// Output:
	// [0 1 2]
	// [3 4]
}

func TestKeyedSplitIn(t *testing.T) {
	assertPanicsWith(t, collect.ErrInvalidArgument, func() { collect.Of(1).Keyed().SplitIn(0) })
	assertStopsEarly(t, collect.Range(1, 9).Keyed().SplitIn(9).Seq())
}

func TestSplitInMaxInt(t *testing.T) {
	assertChunks(t, collect.Of(1, 2).SplitIn(math.MaxInt), [][]int{{1}, {2}})
}

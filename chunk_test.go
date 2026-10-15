package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Chunk() {
	chunks := collect.Of(1, 2, 3, 4, 5, 6, 7).Chunk(4)

	for chunk := range chunks.Seq() {
		fmt.Println(chunk.All())
	}
	// Output:
	// [1 2 3 4]
	// [5 6 7]
}

// The chunks are a Collection like any other, so they chunk again.
func ExampleCollection_Chunk_nested() {
	quads := collect.Of(1, 2, 3, 4, 5, 6, 7, 8).
		Chunk(2).
		Chunk(2).
		Map(func(c collect.Collection[collect.Collection[int]]) [][]int {
			return c.Map(collect.Collection[int].All).All()
		}).
		All()

	fmt.Println(quads)
	// Output: [[[1 2] [3 4]] [[5 6] [7 8]]]
}

// Each pair keeps its key, as in Laravel: the second chunk starts at key 2.
func ExampleKeyedCollection_Chunk() {
	chunks := collect.Of("a", "b", "c", "d", "e").Keyed().Chunk(2)

	for chunk := range chunks.Seq() {
		fmt.Println(chunk.Keys().All(), chunk.Values().All())
	}
	// Output:
	// [0 1] [a b]
	// [2 3] [c d]
	// [4] [e]
}

func TestChunk(t *testing.T) {
	c := collect.Of(1, 2, 3)
	assertChunks(t, c.Chunk(3), [][]int{{1, 2, 3}})
	assertChunks(t, c.Chunk(9), [][]int{{1, 2, 3}})
	assertChunks(t, collect.Of[int]().Chunk(2), nil)
	if n := collect.Of[int]().Keyed().Chunk(2).Count(); n != 0 {
		t.Errorf("Chunk(2) of nothing has %d chunks, want 0", n)
	}
}

// The chunks are a Collection, so every method applies to them.
func TestChunkThenAnyMethod(t *testing.T) {
	sizes := collect.Of(1, 2, 3, 4, 5).Chunk(2).
		Filter(func(c collect.Collection[int]) bool { return c.Count() == 2 }).
		Keyed().
		Map(func(i int, c collect.Collection[int]) int { return i*10 + c.Count() })
	assertPairs(t, sizes, []int{0, 1}, []int{2, 12})
	pairs := collect.Range(1, 8).Keyed().Chunk(2).Chunk(2).All()
	assertPairs(t, pairs[1].All()[0], []int{4, 5}, []int{5, 6})
	if n := collect.Range(1, 8).Chunk(2).Chunk(2).Chunk(2).Count(); n != 1 {
		t.Errorf("Chunk(2) three times has %d chunks, want 1", n)
	}
}

// A chunk stays the same after the next one is made.
func TestChunkKeepsEachChunk(t *testing.T) {
	all := collect.Of(1, 2, 3, 4).Chunk(2).All()
	assertEqual(t, all[0].All(), []int{1, 2})
	assertEqual(t, all[1].All(), []int{3, 4})
}

func TestChunkPanicsBelowOne(t *testing.T) {
	assertPanics(t, func() { collect.Of(1).Chunk(0) })
	assertPanics(t, func() { collect.Of(1).Keyed().Chunk(-1) })
}

func TestChunkStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Chunk(2).Seq())
	assertStopsEarly(t, endless().Keyed().Chunk(2).Seq())
	for chunk := range collect.Of(1, 2, 3, 4).Keyed().Chunk(4).Seq() {
		for range chunk.Seq() {
			break
		}
	}
}

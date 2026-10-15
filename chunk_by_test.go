package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of chunkBy.
func ExampleCollection_ChunkBy() {
	parity := func(n int) bool { return n%2 == 0 }
	chunks := collect.Of(1, 3, 2, 4, 5).ChunkBy(parity)

	fmt.Println(chunks.Map(collect.Collection[int].All).All())
	// Output: [[1 3] [2 4] [5]]
}

func TestChunkBy(t *testing.T) {
	assertChunks(t, letters("AABBCCCD").ChunkBy(collect.Identity), [][]string{{"A", "A"}, {"B", "B"}, {"C", "C", "C"}, {"D"}})
	assertChunks(t, collect.Of[int]().ChunkBy(collect.Identity), nil)
}

func TestChunkByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().ChunkBy(collect.Identity).Seq())
}

// Each pair of a chunk keeps its key.
func ExampleKeyedCollection_ChunkBy() {
	chunks := collect.Of(1, 3, 2).Keyed().ChunkBy(func(_, n int) bool { return n%2 == 0 })

	for chunk := range chunks.Seq() {
		fmt.Println(chunk.Keys().All())
	}
	// Output:
	// [0 1]
	// [2]
}

func TestKeyedChunkBy(t *testing.T) {
	if n := collect.Of[int]().Keyed().ChunkBy(func(_, v int) int { return v }).Count(); n != 0 {
		t.Errorf("ChunkBy() of nothing has %d chunks", n)
	}
	assertStopsEarly(t, endless().Keyed().ChunkBy(func(k, _ int) int { return k }).Seq())
}

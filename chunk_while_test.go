package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of chunkWhile.
func ExampleCollection_ChunkWhile() {
	same := func(letter string, chunk collect.Collection[string]) bool {
		last, _ := chunk.Last()
		return letter == last
	}
	chunks := collect.Of("A", "A", "B", "B", "C", "C", "C", "D").ChunkWhile(same)

	fmt.Println(chunks.Map(collect.Collection[string].All).All())
	// Output: [[A A] [B B] [C C C] [D]]
}

func TestChunkWhileOfNothing(t *testing.T) {
	assertChunks(t, letters("").ChunkWhile(func(string, collect.Collection[string]) bool { return true }), nil)
}

func TestChunkWhileStopsEarly(t *testing.T) {
	never := func(int, collect.Collection[int]) bool { return false }
	assertStopsEarly(t, endless().ChunkWhile(never).Seq())
}

func ExampleKeyedCollection_ChunkWhile() {
	same := func(_ int, letter string, chunk collect.KeyedCollection[int, string]) bool {
		_, last, _ := chunk.Last()
		return letter == last
	}
	chunks := collect.Of("A", "A", "B").Keyed().ChunkWhile(same)

	for chunk := range chunks.Seq() {
		fmt.Println(chunk.Keys().All(), chunk.Values().All())
	}
	// Output:
	// [0 1] [A A]
	// [2] [B]
}

func TestKeyedChunkWhile(t *testing.T) {
	never := func(int, int, collect.KeyedCollection[int, int]) bool { return false }
	if n := collect.Of[int]().Keyed().ChunkWhile(never).Count(); n != 0 {
		t.Errorf("ChunkWhile() of nothing has %d chunks", n)
	}
	assertStopsEarly(t, endless().Keyed().ChunkWhile(never).Seq())
}

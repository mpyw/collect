package collect_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_FlatMap() {
	words := collect.Of("a b", "c").
		FlatMap(func(s string) collect.Collection[string] { return collect.FromSlice(strings.Fields(s)) }).
		All()

	fmt.Println(words)
	// Output: [a b c]
}

// FlatMap with a function that returns its argument flattens chunks.
func ExampleCollection_FlatMap_flatten() {
	flat := collect.Of(1, 2, 3, 4, 5).
		Chunk(2).
		FlatMap(collect.Identity).
		All()

	fmt.Println(flat)
	// Output: [1 2 3 4 5]
}

func TestFlatMapOfNothing(t *testing.T) {
	none := func(int) collect.Collection[string] { return collect.Collection[string]{} }
	assertEqual(t, collect.Of(1, 2).FlatMap(none).All(), nil)
}

func TestFlatMapStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().FlatMap(func(v int) collect.Collection[int] { return collect.Of(v, v) }).Seq())
}

func ExampleKeyedCollection_FlatMap() {
	c := collect.FromSeq2(slices.All([]string{"a b", "c"})).
		FlatMap(func(_ int, s string) collect.KeyedCollection[int, string] {
			return collect.FromSeq2(slices.All(strings.Fields(s)))
		})

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [0 1 0] [a b c]
}

func TestKeyedFlatMapStopsEarly(t *testing.T) {
	twice := func(k, v int) collect.KeyedCollection[int, int] { return collect.Of(v, v).Keyed() }
	assertStopsEarly2(t, endless().Keyed().FlatMap(twice).Seq())
}

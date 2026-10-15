package collect_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/mpyw/collect"
)

// The order is random, so the example shows what does not change.
func ExampleCollection_Shuffle() {
	shuffled := collect.Range(1, 5).Shuffle().All()
	slices.Sort(shuffled)

	fmt.Println(shuffled)
	// Output: [1 2 3 4 5]
}

func TestShuffle(t *testing.T) {
	got := collect.Range(1, 50).Shuffle().All()
	if slices.Equal(got, collect.Range(1, 50).All()) {
		t.Error("Shuffle kept the order of 50 values")
	}
	slices.Sort(got)
	assertEqual(t, got, collect.Range(1, 50).All())
}

func TestShuffleStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Range(1, 9).Shuffle().Seq())
}

// Each pair keeps its key.
func ExampleKeyedCollection_Shuffle() {
	c := collect.Of(10, 20, 30).Keyed().Shuffle()

	fmt.Println(c.Every(func(i, n int) bool { return n == (i+1)*10 }), c.Count())
	// Output: true 3
}

func TestKeyedShuffleStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().Shuffle().Seq())
}

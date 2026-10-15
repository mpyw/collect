package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of diffKeys.
func ExampleKeyedCollection_DiffKeys() {
	c := keyedOf([]string{"one", "two", "three", "four", "five"}, []int{10, 20, 30, 40, 50})
	got := c.DiffKeys(keyedOf([]string{"two", "four", "six"}, []int{2, 4, 6}))

	fmt.Println(got.Keys().All(), got.Values().All())
	// Output: [one three five] [10 30 50]
}

func TestDiffByKeysStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().DiffKeys(collect.Of(1).Keyed()).Seq())
}

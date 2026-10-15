package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of replace.
func ExampleKeyedCollection_Replace() {
	c := collect.Of("Taylor", "Abigail", "James").Keyed().
		Replace(keyedOf([]int{1, 3}, []string{"Victoria", "Finn"}))

	fmt.Println(c.Values().All())
	// Output: [Taylor Victoria James Finn]
}

// A key of other that comes twice counts once, with its last value.
func TestReplaceRepeatedKey(t *testing.T) {
	got := collect.Of("Taylor").Keyed().Replace(keyedOf([]int{5, 5}, []string{"x", "y"}))
	assertPairs(t, got, []int{0, 5}, []string{"Taylor", "y"})
}

func TestReplaceStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().Replace(collect.Of(1).Keyed()).Seq())
	assertStopsEarly2(t, collect.Of(1).Keyed().Replace(collect.Range(1, 9).Keyed()).Seq())
}

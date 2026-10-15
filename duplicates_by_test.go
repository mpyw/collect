package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of duplicates.
func ExampleCollection_DuplicatesBy() {
	fmt.Println(collect.Of("a", "b", "a", "c", "b").DuplicatesBy(collect.Identity).All())
	// Output: [a b]
}

func TestDuplicatesByEachTime(t *testing.T) {
	assertEqual(t, collect.Of(1, 1, 1).DuplicatesBy(collect.Identity).All(), []int{1, 1})
}

func TestDuplicatesByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().DuplicatesBy(func(v int) int { return v % 2 }).Seq())
}

// Laravel's example of duplicates: each repeat, with its own key.
func ExampleKeyedCollection_DuplicatesBy() {
	c := collect.Of("a", "b", "a", "c", "b").Keyed().DuplicatesBy(func(_ int, s string) string { return s })

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [2 4] [a b]
}

func TestKeyedDuplicatesByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().DuplicatesBy(func(_, v int) int { return v % 2 }).Seq())
}

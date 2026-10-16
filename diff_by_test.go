package collect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of diff.
func ExampleCollection_DiffBy() {
	fmt.Println(collect.Of(1, 2, 3, 4, 5).DiffBy(collect.Of(2, 4, 6, 8), collect.Identity).All())
	// Output: [1 3 5]
}

func TestDiffByKeyFunction(t *testing.T) {
	assertEqual(t, collect.Of("A", "b").DiffBy(collect.Of("a"), strings.ToLower).All(), []string{"b"})
}

func TestDiffByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().DiffBy(collect.Of(2), collect.Identity).Seq())
}

// DiffBy compares the values, and keeps the keys, as in Laravel.
func ExampleKeyedCollection_DiffBy() {
	c := collect.Of(1, 2, 3, 4, 5).Keyed().DiffBy(collect.Of(2, 4).Keyed(), collect.Identity)

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [0 2 4] [1 3 5]
}

func TestKeyedDiffByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().DiffBy(collect.Of(2).Keyed(), collect.Identity).Seq())
}

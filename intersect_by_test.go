package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of intersect.
func ExampleCollection_IntersectBy() {
	got := collect.Of("Desk", "Sofa", "Chair").IntersectBy(collect.Of("Desk", "Chair", "Bookcase"), collect.Identity)

	fmt.Println(got.All())
	// Output: [Desk Chair]
}

func TestIntersectByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().IntersectBy(collect.Range(1, 9), collect.Identity).Seq())
}

func ExampleKeyedCollection_IntersectBy() {
	c := collect.Of("Desk", "Sofa", "Chair").Keyed().IntersectBy(collect.Of("Desk", "Chair").Keyed(), collect.Identity)

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [0 2] [Desk Chair]
}

func TestKeyedIntersectByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().IntersectBy(collect.Range(1, 9).Keyed(), collect.Identity).Seq())
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of intersectAssoc.
func ExampleKeyedCollection_IntersectAssocBy() {
	c := keyedOf([]string{"color", "size", "material"}, []string{"red", "M", "cotton"})
	got := c.IntersectAssocBy(keyedOf([]string{"color", "size", "material"}, []string{"blue", "M", "polyester"}), collect.Identity)

	fmt.Println(got.All())
	// Output: map[size:M]
}

func TestIntersectByAssocStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().IntersectAssocBy(endless().Keyed().Take(9), collect.Identity).Seq())
}

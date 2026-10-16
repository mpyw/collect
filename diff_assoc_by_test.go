package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of diffAssoc.
func ExampleKeyedCollection_DiffAssocBy() {
	c := keyedOf([]string{"color", "type", "remain"}, []any{"orange", "fruit", 6})
	got := c.DiffAssocBy(keyedOf([]string{"color", "type", "used"}, []any{"yellow", "fruit", 6}), collect.Identity)

	fmt.Println(got.ToMap())
	// Output: map[color:orange remain:6]
}

func TestDiffByAssocStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().DiffAssocBy(collect.Of(1).Keyed(), collect.Identity).Seq())
}

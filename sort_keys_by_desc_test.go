package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's sortKeysDesc.
func ExampleKeyedCollection_SortKeysByDesc() {
	c := keyedOf([]string{"id", "first", "last"}, []any{22345, "John", "Doe"}).SortKeysByDesc(collect.Identity)

	fmt.Println(c.Keys().All())
	// Output: [last id first]
}

func TestSortKeysByDescStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().SortKeysByDesc(collect.Identity).Seq())
}

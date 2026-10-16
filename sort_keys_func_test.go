package collect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of sortKeysUsing, with the order of strnatcasecmp's
// letters.
func ExampleKeyedCollection_SortKeysFunc() {
	fold := func(a, b string) int { return strings.Compare(strings.ToLower(a), strings.ToLower(b)) }
	c := keyedOf([]string{"ID", "first", "last"}, []any{22345, "John", "Doe"}).SortKeysFunc(fold)

	fmt.Println(c.Keys().All())
	// Output: [first ID last]
}

func TestSortKeysFuncStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().SortKeysFunc(func(a, b int) int { return b - a }).Seq())
}

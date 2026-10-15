package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

// Laravel's example of sortKeys.
func ExampleKeyedCollection_SortKeysBy() {
	c := keyedOf([]string{"id", "first", "last"}, []any{22345, "John", "Doe"}).SortKeysBy(collect.Identity)

	fmt.Println(c.Keys().All())
	// Output: [first id last]
}

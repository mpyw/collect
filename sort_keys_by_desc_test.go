package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

// Laravel's sortKeysDesc.
func ExampleKeyedCollection_SortKeysByDesc() {
	c := keyedOf([]string{"id", "first", "last"}, []any{22345, "John", "Doe"}).SortKeysByDesc(collect.Identity)

	fmt.Println(c.Keys().All())
	// Output: [last id first]
}

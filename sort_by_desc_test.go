package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

// Laravel's sortDesc and sortByDesc.
func ExampleCollection_SortByDesc() {
	fmt.Println(collect.Of(5, 3, 1, 2, 4).SortByDesc(collect.Identity).All())
	// Output: [5 4 3 2 1]
}

func ExampleKeyedCollection_SortByDesc() {
	c := collect.Of(3, 1, 2).Keyed().SortByDesc(func(_, n int) int { return n })

	fmt.Println(c.Keys().All())
	// Output: [0 2 1]
}

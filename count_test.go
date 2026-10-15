package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

func ExampleCollection_Count() {
	fmt.Println(collect.Of(1, 2, 3, 4).Count())
	// Output: 4
}

// A key that comes more than once is counted each time.
func ExampleKeyedCollection_Count() {
	c := collect.Of("apple", "avocado").KeyBy(func(s string) string { return s[:1] })

	fmt.Println(c.Count(), len(c.All()))
	// Output: 2 1
}

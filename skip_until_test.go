package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

// Laravel's example of skipUntil.
func ExampleCollection_SkipUntil() {
	fmt.Println(collect.Of(1, 2, 3, 4).SkipUntil(func(n int) bool { return n >= 3 }).All())
	// Output: [3 4]
}

func ExampleKeyedCollection_SkipUntil() {
	c := collect.Of(1, 2, 3, 1).Keyed().SkipUntil(func(_, n int) bool { return n == 3 })

	fmt.Println(c.Keys().All())
	// Output: [2 3]
}

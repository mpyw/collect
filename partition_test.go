package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

// Laravel's example of partition.
func ExampleCollection_Partition() {
	under3, over3 := collect.Of(1, 2, 3, 4, 5, 6).Partition(func(n int) bool { return n < 3 })

	fmt.Println(under3.All(), over3.All())
	// Output: [1 2] [3 4 5 6]
}

func ExampleKeyedCollection_Partition() {
	even, odd := collect.Of("a", "b", "c").Keyed().Partition(func(i int, _ string) bool { return i%2 == 0 })

	fmt.Println(even.Keys().All(), odd.Keys().All())
	// Output: [0 2] [1]
}

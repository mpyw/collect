package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

func ExampleCollection_Pipe() {
	total := collect.Of(1, 2, 3).Pipe(func(c collect.Collection[int]) int { return c.SumBy(collect.Identity) })

	fmt.Println(total)
	// Output: 6
}

func ExampleKeyedCollection_Pipe() {
	fmt.Println(collect.Of("a").Keyed().Pipe(collect.KeyedCollection[int, string].Count))
	// Output: 1
}

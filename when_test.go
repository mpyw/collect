package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

// Laravel's example of when.
func ExampleCollection_When() {
	addFour := func(c collect.Collection[int]) collect.Collection[int] { return c.Concat(collect.Of(4)) }

	fmt.Println(collect.Of(1, 2, 3).When(true, addFour).All())
	fmt.Println(collect.Of(1, 2, 3).When(false, addFour).All())
	// Output:
	// [1 2 3 4]
	// [1 2 3]
}

func ExampleKeyedCollection_When() {
	first := func(c collect.KeyedCollection[int, string]) collect.KeyedCollection[int, string] { return c.Take(1) }

	fmt.Println(collect.Of("a", "b").Keyed().When(true, first).Values().All())
	fmt.Println(collect.Of("a", "b").Keyed().When(false, first).Values().All())
	// Output:
	// [a]
	// [a b]
}

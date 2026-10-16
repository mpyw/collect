package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of unless.
func ExampleCollection_Unless() {
	addFour := func(c collect.Collection[int]) collect.Collection[int] { return c.Concat(collect.Of(4)) }

	fmt.Println(collect.Of(1, 2, 3).Unless(true, addFour).All())
	fmt.Println(collect.Of(1, 2, 3).Unless(false, addFour).All())
	// Output:
	// [1 2 3]
	// [1 2 3 4]
}

func ExampleKeyedCollection_Unless() {
	first := func(c collect.KeyedCollection[int, string]) collect.KeyedCollection[int, string] { return c.Take(1) }

	fmt.Println(collect.Of("a", "b").Keyed().Unless(false, first).Values().All())
	// Output: [a]
}

func TestUnlessStopsEarly(t *testing.T) {
	id := func(c collect.Collection[int]) collect.Collection[int] { return c }
	assertStopsEarly(t, endless().Unless(false, id).Seq())
	assertStopsEarly2(t, endless().Keyed().Unless(false, func(c collect.KeyedCollection[int, int]) collect.KeyedCollection[int, int] { return c }).Seq())
}

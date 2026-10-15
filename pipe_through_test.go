package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_PipeThrough() {
	more := func(c collect.Collection[int]) collect.Collection[int] { return c.Concat(collect.Of(4, 5)) }
	double := func(c collect.Collection[int]) collect.Collection[int] {
		return c.Map(func(n int) int { return n * 2 })
	}

	fmt.Println(collect.Of(1, 2, 3).PipeThrough(more, double).All())
	// Output: [2 4 6 8 10]
}

func ExampleKeyedCollection_PipeThrough() {
	first := func(c collect.KeyedCollection[int, string]) collect.KeyedCollection[int, string] { return c.Take(1) }

	fmt.Println(collect.Of("a", "b").Keyed().PipeThrough(first).Values().All())
	// Output: [a]
}

func TestPipeThroughNothing(t *testing.T) {
	assertEqual(t, collect.Of(1, 2).PipeThrough().All(), []int{1, 2})
}

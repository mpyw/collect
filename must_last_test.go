package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_MustLast() {
	fmt.Println(collect.Of(1, 2, 3).MustLast())
	// Output: 3
}

func ExampleKeyedCollection_MustLast() {
	fmt.Println(collect.Of("a", "b").Keyed().MustLast())
	// Output: 1 b
}

func TestMustLast(t *testing.T) {
	assertPanicsWith(t, collect.ErrItemNotFound, func() { collect.Of[int]().MustLast() })
	assertPanicsWith(t, collect.ErrItemNotFound, func() { collect.FromMap(map[int]int{}).MustLast() })
}

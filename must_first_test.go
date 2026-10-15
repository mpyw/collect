package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_MustFirst() {
	fmt.Println(collect.Of(1, 2, 3).MustFirst())

	defer func() { fmt.Println("panic:", recover()) }()
	collect.Of[int]().MustFirst()
	// Output:
	// 1
	// panic: collect: item not found
}

func ExampleKeyedCollection_MustFirst() {
	fmt.Println(collect.Of("a", "b").Keyed().MustFirst())
	// Output: 0 a
}

func TestMustFirst(t *testing.T) {
	if v := endless().MustFirst(); v != 1 {
		t.Errorf("MustFirst() of an endless source = %v", v)
	}
	assertPanicsWith(t, collect.ErrItemNotFound, func() { collect.FromMap(map[int]int{}).MustFirst() })
}

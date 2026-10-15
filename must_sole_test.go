package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_MustSole() {
	fmt.Println(collect.Of(1, 2, 3).Filter(func(n int) bool { return n == 2 }).MustSole())

	defer func() { fmt.Println("panic:", recover()) }()
	collect.Of(1, 2, 3).MustSole()
	// Output:
	// 2
	// panic: collect: multiple items found
}

func TestMustSole(t *testing.T) {
	assertPanicsWith(t, collect.ErrItemNotFound, func() { collect.Of[int]().MustSole() })
	assertPanicsWith(t, collect.ErrMultipleItemsFound, func() { endless().MustSole() })
}

func ExampleKeyedCollection_MustSole() {
	fmt.Println(collect.FromMap(map[string]int{"only": 1}).MustSole())
	// Output: only 1
}

func TestKeyedMustSole(t *testing.T) {
	assertPanicsWith(t, collect.ErrMultipleItemsFound, func() { endless().Keyed().MustSole() })
}

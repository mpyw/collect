package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleKeyedCollection_MustGet() {
	c := collect.FromMap(map[string]string{"name": "taylor"})
	fmt.Println(c.MustGet("name"))

	defer func() { fmt.Println("panic:", recover()) }()
	c.MustGet("age")
	// Output:
	// taylor
	// panic: collect: item not found: no key age
}

func TestMustGet(t *testing.T) {
	assertPanicsWith(t, collect.ErrItemNotFound, func() { collect.Of(1).Keyed().MustGet(5) })
}

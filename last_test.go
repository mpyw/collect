package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Last() {
	v, ok := collect.Of(1, 2, 3, 4).Last()

	fmt.Println(v, ok)
	// Output: 4 true
}

func ExampleKeyedCollection_Last() {
	k, v, ok := collect.Of("a", "b").Keyed().Last()

	fmt.Println(k, v, ok)
	// Output: 1 b true
}

func TestLastOfNothing(t *testing.T) {
	if _, ok := collect.Of[int]().Last(); ok {
		t.Error("Last() of nothing is ok")
	}
	if _, _, ok := collect.FromMap(map[int]int{}).Last(); ok {
		t.Error("Last() of nothing is ok")
	}
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// A key that comes more than once keeps its last value.
func ExampleKeyedCollection_ToMap() {
	byInitial := collect.Of("apple", "avocado", "banana").
		KeyBy(func(s string) string { return s[:1] }).
		ToMap()

	fmt.Println(byInitial)
	// Output: map[a:avocado b:banana]
}

func TestToMapOfNothing(t *testing.T) {
	if m := collect.Of[int]().Keyed().ToMap(); m == nil || len(m) != 0 {
		t.Errorf("ToMap() = %#v, want an empty map", m)
	}
}

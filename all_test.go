package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_All() {
	fmt.Println(collect.Of(1, 2, 3).All())
	// Output: [1 2 3]
}

// A key that comes more than once keeps its last value.
func ExampleKeyedCollection_All() {
	byInitial := collect.Of("apple", "avocado", "banana").
		KeyBy(func(s string) string { return s[:1] }).
		All()

	fmt.Println(byInitial)
	// Output: map[a:avocado b:banana]
}

func TestAllOfNothing(t *testing.T) {
	assertEqual(t, collect.Of[int]().All(), nil)
	if m := collect.Of[int]().Keyed().All(); m == nil || len(m) != 0 {
		t.Errorf("All() = %#v, want an empty map", m)
	}
}

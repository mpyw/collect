package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of takeUntil.
func ExampleCollection_TakeUntil() {
	fmt.Println(collect.Of(1, 2, 3, 4).TakeUntil(func(n int) bool { return n >= 3 }).All())
	// Output: [1 2]
}

func ExampleKeyedCollection_TakeUntil() {
	fmt.Println(collect.Of("a", "b", "c").Keyed().TakeUntil(func(i int, _ string) bool { return i == 2 }).Values().All())
	// Output: [a b]
}

func TestTakeUntilStopsEarly(t *testing.T) {
	never := func(int) bool { return false }
	assertStopsEarly(t, endless().TakeUntil(never).Seq())
	assertStopsEarly2(t, endless().Keyed().TakeUntil(func(int, int) bool { return false }).Seq())
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of takeWhile.
func ExampleCollection_TakeWhile() {
	fmt.Println(collect.Of(1, 2, 3, 4).TakeWhile(func(n int) bool { return n < 3 }).All())
	// Output: [1 2]
}

func ExampleKeyedCollection_TakeWhile() {
	fmt.Println(collect.Of("a", "b", "c").Keyed().TakeWhile(func(i int, _ string) bool { return i < 1 }).Values().All())
	// Output: [a]
}

func TestTakeWhileStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().TakeWhile(func(int) bool { return true }).Seq())
	assertStopsEarly2(t, endless().Keyed().TakeWhile(func(int, int) bool { return true }).Seq())
}

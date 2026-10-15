package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of skipWhile.
func ExampleCollection_SkipWhile() {
	fmt.Println(collect.Of(1, 2, 3, 4).SkipWhile(func(n int) bool { return n <= 3 }).All())
	// Output: [4]
}

func ExampleKeyedCollection_SkipWhile() {
	c := collect.Of(1, 2, 3, 1).Keyed().SkipWhile(func(_, n int) bool { return n < 3 })

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [2 3] [3 1]
}

func TestSkipWhileStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().SkipWhile(func(v int) bool { return v < 3 }).Seq())
	assertStopsEarly2(t, endless().Keyed().SkipWhile(func(k, _ int) bool { return k < 3 }).Seq())
}

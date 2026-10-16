package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of skipUntil.
func ExampleCollection_SkipUntil() {
	fmt.Println(collect.Of(1, 2, 3, 4).SkipUntil(func(n int) bool { return n >= 3 }).All())
	// Output: [3 4]
}

func ExampleKeyedCollection_SkipUntil() {
	c := collect.Of(1, 2, 3, 1).Keyed().SkipUntil(func(_, n int) bool { return n == 3 })

	fmt.Println(c.Keys().All())
	// Output: [2 3]
}

func TestSkipUntilStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().SkipUntil(func(v int) bool { return v > 2 }).Seq())
	assertStopsEarly2(t, endless().Keyed().SkipUntil(func(k, _ int) bool { return k > 2 }).Seq())
}

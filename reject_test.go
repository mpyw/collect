package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Reject() {
	small := collect.Of(1, 2, 3, 4).
		Reject(func(n int) bool { return n > 2 }).
		All()

	fmt.Println(small)
	// Output: [1 2]
}

func ExampleKeyedCollection_Reject() {
	c := collect.Of("a", "b", "c").Keyed().
		Reject(func(i int, _ string) bool { return i == 0 })

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [1 2] [b c]
}

func TestRejectStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Reject(func(int) bool { return false }).Seq())
	assertStopsEarly2(t, endless().Keyed().Reject(func(int, int) bool { return false }).Seq())
}

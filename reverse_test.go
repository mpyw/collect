package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of reverse.
func ExampleCollection_Reverse() {
	fmt.Println(collect.Of("a", "b", "c", "d", "e").Reverse().All())
	// Output: [e d c b a]
}

func ExampleKeyedCollection_Reverse() {
	fmt.Println(collect.Of("a", "b").Keyed().Reverse().Keys().All())
	// Output: [1 0]
}

func TestReverseStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Range(1, 9).Reverse().Seq())
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().Reverse().Seq())
}

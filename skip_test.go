package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Skip() {
	fmt.Println(collect.Range(1, 10).Skip(4).All())
	// Output: [5 6 7 8 9 10]
}

func ExampleKeyedCollection_Skip() {
	fmt.Println(collect.Of("a", "b", "c").Keyed().Skip(2).Keys().All())
	// Output: [2]
}

func TestSkip(t *testing.T) {
	c := collect.Of(1, 2, 3)
	for _, n := range []int{-1, 0} {
		assertEqual(t, c.Skip(n).All(), []int{1, 2, 3})
		assertEqual(t, c.Keyed().Skip(n).Values().All(), []int{1, 2, 3})
	}
	assertEqual(t, c.Skip(5).All(), nil)
}

func TestSkipStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Skip(1).Seq())
	assertStopsEarly2(t, endless().Keyed().Skip(1).Seq())
}

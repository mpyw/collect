package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of nth.
func ExampleCollection_Nth() {
	c := collect.Of("a", "b", "c", "d", "e", "f")

	fmt.Println(c.Nth(4, 0).All())
	fmt.Println(c.Nth(4, 1).All())
	// Output:
	// [a e]
	// [b f]
}

func TestNth(t *testing.T) {
	c := letters("abcdef")
	assertEqual(t, c.Nth(2, -3).All(), []string{"d", "f"})
	assertPanics(t, func() { c.Nth(0, 0) })
}

func TestNthStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Nth(2, 1).Seq())
}

func ExampleKeyedCollection_Nth() {
	fmt.Println(collect.Of("a", "b", "c", "d", "e", "f").Keyed().Nth(4, 1).Keys().All())
	// Output: [1 5]
}

func TestKeyedNth(t *testing.T) {
	assertEqual(t, letters("abcdef").Keyed().Nth(2, -3).Keys().All(), []int{3, 5})
	assertPanics(t, func() { collect.Of(1).Keyed().Nth(0, 0) })
	assertStopsEarly2(t, endless().Keyed().Nth(2, 1).Seq())
}

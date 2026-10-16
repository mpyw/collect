package collect_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of mode.
func ExampleCollection_ModeBy() {
	type score struct{ Foo int }
	fmt.Println(collect.Of(score{10}, score{10}, score{20}, score{40}).ModeBy(func(s score) int { return s.Foo }).All())
	fmt.Println(collect.Of(1, 1, 2, 4).ModeBy(collect.Identity).All())
	fmt.Println(collect.Of(1, 1, 2, 2).ModeBy(collect.Identity).All())
	// Output:
	// [10]
	// [1]
	// [1 2]
}

func TestModeByOfNothing(t *testing.T) {
	assertEqual(t, collect.Of[int]().ModeBy(collect.Identity).All(), nil)
}

func ExampleKeyedCollection_ModeBy() {
	c := collect.Of(1, 1, 2).Keyed()

	fmt.Println(c.ModeBy(func(_, n int) int { return n }).All())
	// Output: [1]
}

func TestModeByStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Of(1, 2, 3, 4).ModeBy(collect.Identity).Seq())
	assertStopsEarly(t, collect.Of(1, 2, 3, 4).Keyed().ModeBy(collect.Value).Seq())
}

func TestModeByNaN(t *testing.T) {
	if n := collect.Of(math.NaN(), math.NaN()).ModeBy(collect.Identity).Count(); n != 2 {
		t.Errorf("ModeBy of two NaNs has %d keys, want 2", n)
	}
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of mode.
func ExampleCollection_ModeBy() {
	type score struct{ Foo int }
	fmt.Println(collect.Of(score{10}, score{10}, score{20}, score{40}).ModeBy(func(s score) int { return s.Foo }))
	fmt.Println(collect.Of(1, 1, 2, 4).ModeBy(collect.Identity))
	fmt.Println(collect.Of(1, 1, 2, 2).ModeBy(collect.Identity))
	// Output:
	// [10]
	// [1]
	// [1 2]
}

func TestModeByOfNothing(t *testing.T) {
	assertEqual(t, collect.Of[int]().ModeBy(collect.Identity), nil)
}

func ExampleKeyedCollection_ModeBy() {
	c := collect.Of(1, 1, 2).Keyed()

	fmt.Println(c.ModeBy(func(_, n int) int { return n }))
	// Output: [1]
}

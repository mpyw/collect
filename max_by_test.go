package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of max.
func ExampleCollection_MaxBy() {
	type score struct{ Foo int }
	greatest, ok := collect.Of(score{10}, score{20}).MaxBy(func(s score) int { return s.Foo })
	fmt.Println(greatest, ok)

	greatest, ok = collect.Of(1, 2, 3, 4, 5).MaxBy(collect.Identity)
	fmt.Println(greatest, ok)
	// Output:
	// 20 true
	// 5 true
}

func TestMaxByOfNothing(t *testing.T) {
	if _, ok := collect.Of[int]().MaxBy(collect.Identity); ok {
		t.Error("Max() of nothing is ok")
	}
}

func ExampleKeyedCollection_MaxBy() {
	prices := collect.FromMap(map[string]int{"desk": 200, "chair": 100})

	fmt.Println(prices.MaxBy(func(name string, _ int) string { return name }))
	// Output: desk true
}

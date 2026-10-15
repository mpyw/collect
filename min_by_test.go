package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of min.
func ExampleCollection_MinBy() {
	type score struct{ Foo int }
	least, ok := collect.Of(score{10}, score{20}).MinBy(func(s score) int { return s.Foo })
	fmt.Println(least, ok)

	least, ok = collect.Of(1, 2, 3, 4, 5).MinBy(collect.Identity)
	fmt.Println(least, ok)
	// Output:
	// 10 true
	// 1 true
}

func TestMinBy(t *testing.T) {
	if got, _ := collect.Of("b", "a").MinBy(collect.Identity); got != "a" {
		t.Errorf("Min() of strings = %v", got)
	}
	if _, ok := collect.Of[int]().MinBy(collect.Identity); ok {
		t.Error("Min() of nothing is ok")
	}
}

func ExampleKeyedCollection_MinBy() {
	prices := collect.FromMap(map[string]int{"desk": 200, "chair": 100})

	fmt.Println(prices.MinBy(func(_ string, price int) int { return price }))
	// Output: 100 true
}

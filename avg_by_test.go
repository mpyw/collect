package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of avg.
func ExampleCollection_AvgBy() {
	avg, ok := collect.Of(1, 1, 2, 4).AvgBy(collect.Identity)
	fmt.Println(avg, ok)

	type score struct{ Foo int }
	avg, ok = collect.Of(score{10}, score{10}, score{20}, score{40}).AvgBy(func(s score) int { return s.Foo })
	fmt.Println(avg, ok)
	// Output:
	// 2 true
	// 20 true
}

func TestAvgByOfNothing(t *testing.T) {
	if _, ok := collect.Of[int]().AvgBy(collect.Identity); ok {
		t.Error("Avg() of nothing is ok")
	}
}

func ExampleKeyedCollection_AvgBy() {
	prices := collect.FromMap(map[string]int{"desk": 200, "chair": 100})

	fmt.Println(prices.AvgBy(func(_ string, price int) int { return price }))
	// Output: 150 true
}

func TestKeyedAvgByOfNothing(t *testing.T) {
	if _, ok := collect.Of[int]().Keyed().AvgBy(func(_, v int) int { return v }); ok {
		t.Error("Avg() of nothing is ok")
	}
}

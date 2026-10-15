package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of median.
func ExampleCollection_MedianBy() {
	type score struct{ Foo int }
	median, ok := collect.Of(score{10}, score{10}, score{20}, score{40}).MedianBy(func(s score) int { return s.Foo })
	fmt.Println(median, ok)

	median, ok = collect.Of(1, 1, 2, 4).MedianBy(collect.Identity)
	fmt.Println(median, ok)
	// Output:
	// 15 true
	// 1.5 true
}

func TestMedianBy(t *testing.T) {
	if got, _ := collect.Of(10, 30, 20).MedianBy(collect.Identity); got != 20 {
		t.Errorf("Median() of 3 = %v", got)
	}
	if _, ok := collect.Of[int]().MedianBy(collect.Identity); ok {
		t.Error("Median() of nothing is ok")
	}
}

func ExampleKeyedCollection_MedianBy() {
	c := collect.Of(10, 30, 20).Keyed()

	fmt.Println(c.MedianBy(func(_, n int) int { return n }))
	// Output: 20 true
}

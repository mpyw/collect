package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of percentage.
func ExampleCollection_Percentage() {
	c := collect.Of(1, 1, 2, 2, 2, 3)
	one := func(n int) bool { return n == 1 }

	fmt.Println(c.Percentage(one, 2))
	fmt.Println(c.Percentage(one, 3))
	// Output:
	// 33.33 true
	// 33.333 true
}

func TestPercentageOfNothing(t *testing.T) {
	if _, ok := collect.Of[int]().Percentage(func(int) bool { return true }, 2); ok {
		t.Error("Percentage() of nothing is ok")
	}
}

func ExampleKeyedCollection_Percentage() {
	c := collect.Of(1, 1, 2, 2, 2, 3).Keyed()

	fmt.Println(c.Percentage(func(_, n int) bool { return n == 1 }, 2))
	// Output: 33.33 true
}

// A precision past what a float64 holds rounds nothing, rather than give
// NaN.
func TestPercentagePrecision(t *testing.T) {
	odd := func(v int) bool { return v%2 == 1 }
	if got, _ := collect.Of(1, 2, 3).Percentage(odd, 400); got != 200.0/3 {
		t.Errorf("Percentage(400) = %v", got)
	}
	if got, _ := collect.Of(1, 2, 3).Percentage(odd, -400); got != 0 {
		t.Errorf("Percentage(-400) = %v", got)
	}
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Every() {
	fmt.Println(collect.Of(1, 2, 3, 4).Every(func(n int) bool { return n > 2 }))
	fmt.Println(collect.Of[int]().Every(func(n int) bool { return n > 2 }))
	// Output:
	// false
	// true
}

func ExampleKeyedCollection_Every() {
	c := collect.Of(1, 2).Keyed()

	fmt.Println(c.Every(func(i, n int) bool { return n == i+1 }))
	// Output: true
}

func TestEveryStopsReading(t *testing.T) {
	if endless().Every(func(v int) bool { return v < 5 }) {
		t.Error("Every(< 5) of an endless source is true")
	}
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Contains() {
	c := collect.Of(1, 2, 3, 4, 5)

	fmt.Println(c.Contains(func(n int) bool { return n > 5 }))
	fmt.Println(c.Contains(func(n int) bool { return n == 3 }))
	// Output:
	// false
	// true
}

func ExampleKeyedCollection_Contains() {
	c := collect.FromMap(map[string]string{"name": "Desk", "price": "100"})

	fmt.Println(c.Contains(func(_, v string) bool { return v == "Desk" }))
	// Output: true
}

func TestContainsStopsReading(t *testing.T) {
	if !endless().Contains(func(v int) bool { return v == 5 }) {
		t.Error("Contains(5) of an endless source is false")
	}
	if !endless().Keyed().Contains(func(k, _ int) bool { return k == 5 }) {
		t.Error("Contains(key 5) of an endless source is false")
	}
}

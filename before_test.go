package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of before.
func ExampleCollection_Before() {
	v, ok := collect.Of(1, 2, 3, 4, 5).Before(func(n int) bool { return n == 3 })

	fmt.Println(v, ok)
	// Output: 2 true
}

func TestBefore(t *testing.T) {
	is := func(want int) func(int) bool { return func(v int) bool { return v == want } }
	c := collect.Of(1, 2, 3)
	if _, ok := c.Before(is(1)); ok {
		t.Error("Before(1) is ok")
	}
	if _, ok := c.Before(is(9)); ok {
		t.Error("Before(9) is ok")
	}
}

func ExampleKeyedCollection_Before() {
	k, v, ok := collect.Of("a", "b", "c").Keyed().Before(func(_ int, s string) bool { return s == "c" })

	fmt.Println(k, v, ok)
	// Output: 1 b true
}

func TestKeyedBefore(t *testing.T) {
	c := collect.Of("a", "b").Keyed()
	if _, _, ok := c.Before(func(_ int, s string) bool { return s == "a" }); ok {
		t.Error("Before(a) is ok")
	}
	if _, _, ok := c.Before(func(_ int, s string) bool { return s == "z" }); ok {
		t.Error("Before(z) is ok")
	}
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of after.
func ExampleCollection_After() {
	v, ok := collect.Of(1, 2, 3, 4, 5).After(func(n int) bool { return n == 3 })

	fmt.Println(v, ok)
	// Output: 4 true
}

func TestAfter(t *testing.T) {
	is := func(want int) func(int) bool { return func(v int) bool { return v == want } }
	if _, ok := collect.Of(1, 2, 3).After(is(3)); ok {
		t.Error("After(3) of the last value is ok")
	}
	if v, ok := endless().After(is(5)); v != 6 || !ok {
		t.Errorf("After(5) of an endless source = %v, %v", v, ok)
	}
}

func ExampleKeyedCollection_After() {
	k, v, ok := collect.Of("a", "b", "c").Keyed().After(func(_ int, s string) bool { return s == "a" })

	fmt.Println(k, v, ok)
	// Output: 1 b true
}

func TestKeyedAfter(t *testing.T) {
	if _, _, ok := collect.Of("a").Keyed().After(func(_ int, s string) bool { return s == "a" }); ok {
		t.Error("After() of the last pair is ok")
	}
}

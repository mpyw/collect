package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_FirstWhere() {
	v, ok := collect.Of(1, 2, 3, 4).FirstWhere(func(n int) bool { return n > 2 })

	fmt.Println(v, ok)
	// Output: 3 true
}

func ExampleKeyedCollection_FirstWhere() {
	k, v, ok := collect.Of("a", "b", "c").Keyed().
		FirstWhere(func(_ int, s string) bool { return s > "a" })

	fmt.Println(k, v, ok)
	// Output: 1 b true
}

func TestFirstWhere(t *testing.T) {
	if _, ok := collect.Of(1, 2).FirstWhere(func(v int) bool { return v > 2 }); ok {
		t.Error("FirstWhere() found a value that does not match")
	}
	if k, _, ok := endless().Keyed().FirstWhere(func(k, _ int) bool { return k == 4 }); k != 4 || !ok {
		t.Errorf("FirstWhere() of an endless source = %v, %v", k, ok)
	}
}

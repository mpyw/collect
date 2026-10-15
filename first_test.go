package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_First() {
	v, ok := collect.Of(1, 2, 3).First()
	fmt.Println(v, ok)

	_, ok = collect.Of[int]().First()
	fmt.Println(ok)
	// Output:
	// 1 true
	// false
}

func ExampleKeyedCollection_First() {
	k, v, ok := collect.Of("a", "b").Keyed().First()

	fmt.Println(k, v, ok)
	// Output: 0 a true
}

func TestFirst(t *testing.T) {
	if v, ok := endless().First(); v != 1 || !ok {
		t.Errorf("First() of an endless source = %v, %v", v, ok)
	}
	if k, v, ok := collect.FromMap(map[string]int{}).First(); k != "" || v != 0 || ok {
		t.Errorf("First() of nothing = %v, %v, %v", k, v, ok)
	}
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_TapEach() {
	firstTwo := collect.Of(1, 2, 3, 4).
		TapEach(func(n int) { fmt.Println("read", n) }).
		Take(2).
		All()

	fmt.Println(firstTwo)
	// Output:
	// read 1
	// read 2
	// [1 2]
}

func ExampleKeyedCollection_TapEach() {
	collect.Of("a").Keyed().TapEach(func(i int, s string) { fmt.Println(i, s) }).ToMap()
	// Output: 0 a
}

func TestTapEach(t *testing.T) {
	var seen []int
	c := collect.Of(1, 2, 3).TapEach(func(v int) { seen = append(seen, v) })
	if seen != nil {
		t.Fatal("TapEach ran before a read")
	}
	assertEqual(t, c.Take(2).All(), []int{1, 2})
	assertEqual(t, seen, []int{1, 2})
	var keys []int
	collect.Of("a", "b").Keyed().TapEach(func(k int, _ string) { keys = append(keys, k) }).ToMap()
	assertEqual(t, keys, []int{0, 1})
}

func TestTapEachStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().TapEach(func(int) {}).Seq())
	assertStopsEarly2(t, endless().Keyed().TapEach(func(int, int) {}).Seq())
}

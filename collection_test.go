package collect_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleOf() {
	values := []int{1, 2, 3}
	c := collect.Of(values...)
	values[0] = 100 // Of copied values, so c does not change.

	fmt.Println(c.All())
	fmt.Println(collect.Of("a", "b").All())
	// Output:
	// [1 2 3]
	// [a b]
}

func ExampleFromSlice() {
	values := []int{1, 2, 3}
	c := collect.FromSlice(values)
	values[0] = 100 // FromSlice did not copy, so the next read sees it.

	fmt.Println(c.All())
	// Output: [100 2 3]
}

func ExampleFromSeq() {
	naturals := func(yield func(int) bool) {
		for i := 1; ; i++ {
			if !yield(i) {
				return
			}
		}
	}

	squares := collect.FromSeq(naturals).
		Map(func(n int) int { return n * n }).
		Take(5).
		All()

	fmt.Println(squares)
	// Output: [1 4 9 16 25]
}

func ExampleFromSeq2() {
	c := collect.FromSeq2(slices.All([]string{"a", "b"}))

	for k, v := range c.Seq() {
		fmt.Println(k, v)
	}
	// Output:
	// 0 a
	// 1 b
}

func ExampleFromMap() {
	c := collect.FromMap(map[string]int{"apple": 3, "banana": 5})

	fmt.Println(c.Values().SumBy(collect.Identity), c.Count())
	// Output: 8 2
}

func ExampleIdentity() {
	c := collect.Of(3, 1, 2)

	fmt.Println(c.SumBy(collect.Identity))
	fmt.Println(c.SortBy(collect.Identity).All())
	// Output:
	// 6
	// [1 2 3]
}

func TestOfIsReusable(t *testing.T) {
	c := collect.Of(1, 2, 3).Map(func(v int) int { return v * 10 })
	assertEqual(t, c.All(), []int{10, 20, 30})
	assertEqual(t, c.All(), []int{10, 20, 30})
}

func TestFromNil(t *testing.T) {
	assertEqual(t, collect.FromSlice[int](nil).All(), nil)
	assertEqual(t, collect.FromSeq[int](nil).All(), nil)
	if n := collect.FromSeq2[int, int](nil).Count(); n != 0 {
		t.Errorf("FromSeq2(nil) has %d pairs, want 0", n)
	}
	if n := collect.FromMap[string, int](nil).Count(); n != 0 {
		t.Errorf("FromMap(nil) has %d pairs, want 0", n)
	}
}

// FromMap does not copy: a read sees the map as it is then.
func TestFromMapReadsWithNoCopy(t *testing.T) {
	m := map[string]int{"a": 1}
	c := collect.FromMap(m)
	m["b"] = 2
	if got := c.All(); !maps.Equal(got, m) {
		t.Errorf("All() = %v, want %v", got, m)
	}
}

func TestZeroValues(t *testing.T) {
	var c collect.Collection[int]
	var k collect.KeyedCollection[string, int]
	if c.Count() != 0 || k.Count() != 0 {
		t.Error("a zero collection is not empty")
	}
}

func ExampleValue() {
	prices := collect.FromMap(map[string]int{"desk": 200, "chair": 100})

	fmt.Println(prices.SumBy(collect.Value))
	// Output: 300
}

func ExampleKey() {
	c := collect.FromMap(map[string]int{"b": 2, "a": 1, "c": 3}).SortBy(collect.Key)

	fmt.Println(c.Keys().All())
	// Output: [a b c]
}

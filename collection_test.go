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
	if got := c.ToMap(); !maps.Equal(got, m) {
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

// Every source that reads the same values each time gives the same result
// to each read of a chain, with no Remember.
func TestReusableSources(t *testing.T) {
	chain := func(c collect.Collection[int]) []int {
		return c.
			Filter(func(v int) bool { return v != 3 }).
			Map(func(v int) int { return v * 10 }).
			SortByDesc(collect.Identity).
			Chunk(2).
			FlatMap(collect.Identity).
			All()
	}
	combined, err := collect.Combine(collect.Of("a", "b", "c", "d"), collect.Of(1, 2, 3, 4))
	if err != nil {
		t.Fatal(err)
	}
	for name, c := range map[string]collect.Collection[int]{
		"Of":        collect.Of(1, 2, 3, 4),
		"FromSlice": collect.FromSlice([]int{1, 2, 3, 4}),
		"FromMap":   collect.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}).Values(),
		"FromSeq":   collect.FromSeq(slices.Values([]int{1, 2, 3, 4})),
		"Range":     collect.Range(1, 4),
		"Times":     collect.Times(4, collect.Identity[int]),
		"Combine":   combined.Values(),
		"Eager":     once(1, 2, 3, 4).Eager(),
		"Remember":  once(1, 2, 3, 4).Remember(),
	} {
		t.Run(name, func(t *testing.T) {
			for range 3 {
				assertEqual(t, chain(c), []int{40, 20, 10})
			}
		})
	}
}

// Each read of a chain reads its source again: nothing is kept between
// reads but by Remember, Eager and Partition.
func TestEachReadReadsTheSource(t *testing.T) {
	starts := 0
	source := collect.FromSeq(func(yield func(int) bool) {
		starts++
		for _, v := range []int{1, 2, 3} {
			if !yield(v) {
				return
			}
		}
	})
	chain := source.Map(func(v int) int { return v * 2 }).Filter(func(v int) bool { return v > 2 })
	chain.All()
	chain.All()
	if starts != 2 {
		t.Errorf("two reads started the source %d times, want 2", starts)
	}
	kept := chain.Remember()
	kept.All()
	kept.All()
	if starts != 3 {
		t.Errorf("two reads of Remember started the source %d more times, want 1", starts-2)
	}
}

// A source that can be read once is empty on its second read, unless
// Remember or Eager keeps it.
func TestSingleUseSource(t *testing.T) {
	c := once(1, 2)
	assertEqual(t, c.All(), []int{1, 2})
	assertEqual(t, c.All(), nil)
	kept := once(1, 2).Remember()
	assertEqual(t, kept.All(), []int{1, 2})
	assertEqual(t, kept.All(), []int{1, 2})
	eager := once(1, 2).Eager()
	assertEqual(t, eager.All(), []int{1, 2})
	assertEqual(t, eager.All(), []int{1, 2})
}

func ExampleFromEntries() {
	entries := collect.Of(collect.E("a", 1), collect.E("b", 2))

	fmt.Println(collect.FromEntries(entries).ToMap())
	// Output: map[a:1 b:2]
}

func TestFromEntriesUndoesEntries(t *testing.T) {
	k := keyedOf([]string{"b", "a", "b"}, []int{1, 2, 3})
	assertPairs(t, collect.FromEntries(k.Entries()), []string{"b", "a", "b"}, []int{1, 2, 3})
	assertStopsEarly2(t, collect.FromEntries(endless().Keyed().Entries()).Seq())
}

func ExampleE() {
	fmt.Println(collect.E("a", 1))
	// Output: {a 1}
}

func ExampleT2() {
	fmt.Println(collect.Of(collect.T2(1, "a"), collect.T2(2, "b")).All())
	// Output: [{1 a} {2 b}]
}

func ExampleT3() {
	fmt.Println(collect.T3(1, "a", true))
	// Output: {1 a true}
}

func ExampleT4() {
	fmt.Println(collect.T4(1, "a", true, 1.5))
	// Output: {1 a true 1.5}
}

func ExampleT5() {
	t := collect.T5(1, "a", true, 1.5, 'x')
	fmt.Println(t.A, t.B, t.C, t.D, string(t.E))
	// Output: 1 a true 1.5 x
}

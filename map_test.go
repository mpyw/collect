package collect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Map() {
	doubled := collect.Of(1, 2, 3, 4, 5).
		Map(func(n int) int { return n * 2 }).
		All()

	fmt.Println(doubled)
	// Output: [2 4 6 8 10]
}

// Map can change the type of the values in the middle of a chain.
func ExampleCollection_Map_newType() {
	labels := collect.Of(1, 2).
		Map(func(n int) string { return fmt.Sprintf("#%d", n) }).
		All()

	fmt.Println(labels)
	// Output: [#1 #2]
}

func ExampleKeyedCollection_Map() {
	c := collect.Of("a", "b").Keyed().
		Map(func(i int, s string) string { return strings.Repeat(s, i+1) })

	for k, v := range c.Seq() {
		fmt.Println(k, v)
	}
	// Output:
	// 0 a
	// 1 bb
}

func TestMapRunsNothingBeforeRead(t *testing.T) {
	seq, n := counted()
	calls := 0
	c := collect.FromSeq(seq).Map(func(v int) int { calls++; return v })
	if *n != 0 || calls != 0 {
		t.Fatalf("read %d values and made %d calls before any read", *n, calls)
	}
	assertEqual(t, c.Take(2).All(), []int{1, 2})
	if calls != 2 {
		t.Errorf("made %d calls for 2 values", calls)
	}
}

func TestMapStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Map(collect.Identity).Seq())
	assertStopsEarly2(t, endless().Keyed().Map(func(_, v int) int { return v }).Seq())
}

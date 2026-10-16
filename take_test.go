package collect_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Take() {
	c := collect.Range(0, 5)

	fmt.Println(c.Take(3).All())
	fmt.Println(c.Take(-2).All()) // The last 2, as in Laravel.
	// Output:
	// [0 1 2]
	// [4 5]
}

func ExampleKeyedCollection_Take() {
	c := collect.Of("a", "b", "c").Keyed()

	fmt.Println(c.Take(2).Keys().All(), c.Take(-1).Keys().All())
	// Output: [0 1] [2]
}

func TestTake(t *testing.T) {
	c := collect.Of(1, 2, 3)
	for _, tt := range []struct {
		n    int
		want []int
	}{
		{-5, []int{1, 2, 3}},
		{0, nil},
		{5, []int{1, 2, 3}},
	} {
		assertEqual(t, c.Take(tt.n).All(), tt.want)
		assertEqual(t, c.Keyed().Take(tt.n).Values().All(), tt.want)
	}
	assertEqual(t, collect.Of[int]().Take(-2).All(), nil)
}

// Take reads no value past the nth, so it ends an endless source.
func TestTakeReadsNoMore(t *testing.T) {
	for _, n := range []int{0, 3} {
		seq, read := counted()
		collect.FromSeq(seq).Take(n).All()
		if *read != n {
			t.Errorf("Take(%d) read %d values", n, *read)
		}
		seq, read = counted()
		collect.FromSeq(seq).Keyed().Take(n).All()
		if *read != n {
			t.Errorf("Keyed().Take(%d) read %d values", n, *read)
		}
	}
}

func TestTakeStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Take(100).Seq())
	assertStopsEarly2(t, endless().Keyed().Take(100).Seq())
	assertStopsEarly(t, collect.Range(1, 9).Take(-5).Seq())
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().Take(-5).Seq())
}

// math.MinInt, whose negation overflows, takes every value.
func TestTakeMinInt(t *testing.T) {
	assertEqual(t, collect.Of(1, 2).Take(math.MinInt).All(), []int{1, 2})
	assertEqual(t, collect.Of(1, 2).Keyed().Take(math.MinInt).Values().All(), []int{1, 2})
}

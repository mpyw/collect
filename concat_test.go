package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of concat.
func ExampleCollection_Concat() {
	names := collect.Of("John Doe").
		Concat(collect.Of("Jane Doe"), collect.Of("Johnny Doe")).
		All()

	fmt.Println(names)
	// Output: [John Doe Jane Doe Johnny Doe]
}

func TestConcatStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Of(1, 2).Concat(endless()).Seq())
}

// The pairs keep their keys, so a key may come twice.
func ExampleKeyedCollection_Concat() {
	c := collect.Of("a").Keyed().Concat(collect.Of("b").Keyed())

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [0 0] [a b]
}

func TestKeyedConcatStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Of(1).Keyed().Concat(endless().Keyed()).Seq())
}

// The others are copied when Concat is called.
func TestConcatCopiesOthers(t *testing.T) {
	others := []collect.Collection[int]{collect.Of(2)}
	c := collect.Of(1).Concat(others...)
	others[0] = collect.Of(3)
	assertEqual(t, c.All(), []int{1, 2})
	keyedOthers := []collect.KeyedCollection[int, int]{collect.Of(2).Keyed()}
	k := collect.Of(1).Keyed().Concat(keyedOthers...)
	keyedOthers[0] = collect.Of(3).Keyed()
	assertEqual(t, k.Values().All(), []int{1, 2})
}

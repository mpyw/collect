package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Entries() {
	for e := range collect.Of("a", "b").Entries().Seq() {
		fmt.Println(e.Key, e.Value)
	}
	// Output:
	// 0 a
	// 1 b
}

// Entries keeps the order and every repeated key, which ToMap does not.
func ExampleKeyedCollection_Entries() {
	c := collect.Of("apple", "avocado", "banana").KeyBy(func(s string) string { return s[:1] })

	fmt.Println(c.Entries().All())
	fmt.Println(c.ToMap())
	// Output:
	// [{a apple} {a avocado} {b banana}]
	// map[a:avocado b:banana]
}

// The entries are a Collection, so every method applies, Keyed and
// Entries again included.
func TestEntriesNest(t *testing.T) {
	deep := collect.Of("a", "b").Keyed().Entries().Keyed().Entries()
	got := deep.Map(func(e collect.Entry[int, collect.Entry[int, string]]) string { return e.Value.Value }).All()
	assertEqual(t, got, []string{"a", "b"})
}

func TestEntriesStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Entries().Seq())
	assertStopsEarly(t, endless().Keyed().Entries().Seq())
}

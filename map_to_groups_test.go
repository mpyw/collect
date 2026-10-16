package collect_test

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of mapToGroups.
func ExampleCollection_MapToGroups() {
	groups := collect.Of(
		employee{"John Doe", "Sales"},
		employee{"Jane Doe", "Sales"},
		employee{"Johnny Doe", "Marketing"},
	).MapToGroups(func(e employee) (string, string) { return e.Department, e.Name })

	for department, names := range groups.Seq() {
		fmt.Println(department, names.All())
	}
	// Output:
	// Sales [John Doe Jane Doe]
	// Marketing [Johnny Doe]
}

type employee struct{ Name, Department string }

func TestMapToGroupsStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Range(1, 9).MapToGroups(func(v int) (int, int) { return v, v }).Seq())
}

func ExampleKeyedCollection_MapToGroups() {
	groups := collect.FromSeq2(slices.All([]string{"apple", "avocado", "banana"})).
		MapToGroups(func(i int, s string) (string, int) { return s[:1], i })

	for initial, indexes := range groups.Seq() {
		fmt.Println(initial, indexes.All())
	}
	// Output:
	// a [0 1]
	// b [2]
}

func TestKeyedMapToGroupsStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().MapToGroups(func(k, v int) (int, int) { return k, v }).Seq())
}

func TestMapToGroupsNaN(t *testing.T) {
	nan := math.NaN()
	groups := collect.Of(1, 2).MapToGroups(func(v int) (float64, int) { return nan, v })
	assertEqual(t, groups.Map(func(_ float64, g collect.Collection[int]) int { return g.Count() }).Values().All(), []int{1, 1})
	keyed := collect.Of(1).Keyed().MapToGroups(func(_, v int) (float64, int) { return nan, v })
	assertEqual(t, keyed.Map(func(_ float64, g collect.Collection[int]) int { return g.Count() }).Values().All(), []int{1})
}

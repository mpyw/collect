package collect_test

import (
	"fmt"
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

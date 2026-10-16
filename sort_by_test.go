package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of sort and sortBy.
func ExampleCollection_SortBy() {
	fmt.Println(collect.Of(5, 3, 1, 2, 4).SortBy(collect.Identity).All())

	type product struct {
		Name  string
		Price int
	}
	products := collect.Of(product{"Desk", 200}, product{"Chair", 100}, product{"Bookcase", 150})
	fmt.Println(products.SortBy(func(p product) int { return p.Price }).All())
	// Output:
	// [1 2 3 4 5]
	// [{Chair 100} {Bookcase 150} {Desk 200}]
}

func ExampleKeyedCollection_SortBy() {
	c := collect.Of(3, 1, 2).Keyed().SortBy(func(_, n int) int { return n })

	fmt.Println(c.Keys().All())
	// Output: [1 2 0]
}

func TestSortByIsStable(t *testing.T) {
	age := func(u user) int { return u.Age }
	users := collect.Of(user{"Alice", 31}, user{"Bob", 19}, user{"Carol", 31}).SortBy(age)
	assertEqual(t, users.Map(func(u user) string { return u.Name }).All(), []string{"Bob", "Alice", "Carol"})
}

func TestSortByStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Range(9, 1).SortBy(collect.Identity).Seq())
	assertStopsEarly2(t, collect.Range(9, 1).Keyed().SortBy(collect.Value).Seq())
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of groupBy.
func ExampleCollection_GroupBy() {
	groups := collect.Of(
		account{"account-x10", "Chair"},
		account{"account-x10", "Bookcase"},
		account{"account-x11", "Desk"},
	).GroupBy(func(a account) string { return a.ID })

	for id, accounts := range groups.Seq() {
		fmt.Println(id, accounts.Map(func(a account) string { return a.Product }).All())
	}
	// Output:
	// account-x10 [Chair Bookcase]
	// account-x11 [Desk]
}

// The groups of a KeyedCollection keep their pairs' keys.
func ExampleKeyedCollection_GroupBy() {
	groups := collect.Of(10, 11, 12, 13).Keyed().
		GroupBy(func(_, n int) string { return map[bool]string{true: "even", false: "odd"}[n%2 == 0] })

	for parity, group := range groups.Seq() {
		fmt.Println(parity, group.Keys().All(), group.Values().All())
	}
	// Output:
	// even [0 2] [10 12]
	// odd [1 3] [11 13]
}

type account struct {
	ID      string
	Product string
}

func TestGroupByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Range(1, 9).GroupBy(collect.Identity).Seq())
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().GroupBy(func(k, _ int) int { return k }).Seq())
	odd, _ := collect.Of(1, 3, 5, 7).Keyed().GroupBy(func(int, int) bool { return true }).Get(true)
	assertStopsEarly2(t, odd.Seq())
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of whereNotIn.
func ExampleCollection_WhereNotInBy() {
	products := collect.Of(
		priced{"Desk", 200},
		priced{"Chair", 100},
		priced{"Bookcase", 150},
		priced{"Door", 100},
	)
	price := func(p priced) int { return p.Price }

	fmt.Println(products.WhereNotInBy(price, 150, 200).All())
	// Output: [{Chair 100} {Door 100}]
}

func TestWhereNotInByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().WhereNotInBy(collect.Identity, 2).Seq())
}

func ExampleKeyedCollection_WhereNotInBy() {
	c := collect.FromMap(map[string]int{"desk": 200, "chair": 100, "door": 150})

	fmt.Println(c.WhereNotInBy(func(_ string, price int) int { return price }, 100, 200).All())
	// Output: map[door:150]
}

func TestKeyedWhereNotInByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().WhereNotInBy(func(k, _ int) int { return k }, -1).Seq())
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of whereIn.
func ExampleCollection_WhereInBy() {
	products := collect.Of(
		priced{"Desk", 200},
		priced{"Chair", 100},
		priced{"Bookcase", 150},
		priced{"Door", 100},
	)
	price := func(p priced) int { return p.Price }

	fmt.Println(products.WhereInBy(price, 150, 200).All())
	// Output: [{Desk 200} {Bookcase 150}]
}

func TestWhereInByNoKeys(t *testing.T) {
	assertEqual(t, collect.Of(1, 2).WhereInBy(collect.Identity).All(), nil)
}

func TestWhereInByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().WhereInBy(func(v int) int { return v % 2 }, 1).Seq())
}

func ExampleKeyedCollection_WhereInBy() {
	c := collect.FromMap(map[string]int{"desk": 200, "chair": 100, "door": 150})

	fmt.Println(c.WhereInBy(func(_ string, price int) int { return price }, 100, 200).ToMap())
	// Output: map[chair:100 desk:200]
}

func TestKeyedWhereInByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().WhereInBy(func(k, _ int) int { return k % 2 }, 1).Seq())
}

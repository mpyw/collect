package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of whereBetween.
func ExampleCollection_WhereBetweenBy() {
	products := collect.Of(
		priced{"Desk", 200},
		priced{"Chair", 80},
		priced{"Bookcase", 150},
		priced{"Pencil", 30},
		priced{"Door", 100},
	)
	price := func(p priced) int { return p.Price }

	fmt.Println(products.WhereBetweenBy(price, 100, 200).All())
	// Output: [{Desk 200} {Bookcase 150} {Door 100}]
}

func TestWhereBetweenByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().WhereBetweenBy(collect.Identity, 1, 100).Seq())
}

func ExampleKeyedCollection_WhereBetweenBy() {
	c := collect.FromMap(map[string]int{"desk": 200, "chair": 80, "door": 100})

	fmt.Println(c.WhereBetweenBy(func(_ string, price int) int { return price }, 100, 200).All())
	// Output: map[desk:200 door:100]
}

func TestKeyedWhereBetweenByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().WhereBetweenBy(func(k, _ int) int { return k }, 0, 100).Seq())
}

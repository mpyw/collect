package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of whereNotBetween.
func ExampleCollection_WhereNotBetweenBy() {
	products := collect.Of(
		priced{"Desk", 200},
		priced{"Chair", 80},
		priced{"Bookcase", 150},
		priced{"Pencil", 30},
		priced{"Door", 100},
	)
	price := func(p priced) int { return p.Price }

	fmt.Println(products.WhereNotBetweenBy(price, 100, 200).All())
	// Output: [{Chair 80} {Pencil 30}]
}

func TestWhereNotBetweenByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().WhereNotBetweenBy(collect.Identity, -2, -1).Seq())
}

func ExampleKeyedCollection_WhereNotBetweenBy() {
	c := collect.FromMap(map[string]int{"desk": 200, "chair": 80, "door": 100})

	fmt.Println(c.WhereNotBetweenBy(func(_ string, price int) int { return price }, 100, 200).All())
	// Output: map[chair:80]
}

func TestKeyedWhereNotBetweenByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().WhereNotBetweenBy(func(k, _ int) int { return k }, -2, -1).Seq())
}

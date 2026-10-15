package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of unique.
func ExampleCollection_UniqueBy() {
	fmt.Println(collect.Of(1, 1, 2, 2, 3, 4, 2).UniqueBy(collect.Identity).All())

	type product struct{ Name, Brand string }
	products := collect.Of(product{"iPhone 6", "Apple"}, product{"iPhone 5", "Apple"}, product{"Galaxy S6", "Samsung"})
	fmt.Println(products.UniqueBy(func(p product) string { return p.Brand }).All())
	// Output:
	// [1 2 3 4]
	// [{iPhone 6 Apple} {Galaxy S6 Samsung}]
}

func TestUniqueByStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().UniqueBy(collect.Identity).Seq())
}

func ExampleKeyedCollection_UniqueBy() {
	c := collect.Of(1, 1, 2).Keyed().UniqueBy(func(_, n int) int { return n })

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [0 2] [1 2]
}

func TestKeyedUniqueByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().UniqueBy(func(k, _ int) int { return k }).Seq())
}

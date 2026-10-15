package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of only.
func ExampleKeyedCollection_Only() {
	c := keyedOf([]string{"product_id", "name", "price", "discount"}, []any{1, "Desk", 100, false})

	fmt.Println(c.Only("product_id", "name").All())
	// Output: map[name:Desk product_id:1]
}

// Only stops once it has each key, so it ends an endless source, and skips
// a key that comes again.
func TestOnly(t *testing.T) {
	assertPairs(t, collect.Of(1).Keyed().Only(), nil, nil)
	assertEqual(t, endless().Keyed().Only(4, 2).Keys().All(), []int{2, 4})
	assertEqual(t, keyedOf([]int{1, 1}, []int{5, 6}).Only(1).Values().All(), []int{5})
}

func TestOnlyStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().Only(1, 2, 3, 4).Seq())
}

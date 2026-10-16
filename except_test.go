package collect_test

import (
	"fmt"
	"testing"
)

// Laravel's example of except.
func ExampleKeyedCollection_Except() {
	c := keyedOf([]string{"product_id", "price", "discount"}, []any{1, 100, false})

	fmt.Println(c.Except("price", "discount").ToMap())
	// Output: map[product_id:1]
}

func TestExceptStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().Except(1).Seq())
}

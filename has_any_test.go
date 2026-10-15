package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of hasAny.
func ExampleKeyedCollection_HasAny() {
	c := collect.FromMap(map[string]int{"account_id": 1, "product": 2, "amount": 5})

	fmt.Println(c.HasAny("product", "price"))
	fmt.Println(c.HasAny("name", "price"))
	// Output:
	// true
	// false
}

func TestHasAnyNoKeys(t *testing.T) {
	if collect.Of(1).Keyed().HasAny() {
		t.Error("HasAny() of no keys is true")
	}
}

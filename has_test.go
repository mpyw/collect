package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of has.
func ExampleKeyedCollection_Has() {
	c := collect.FromMap(map[string]int{"account_id": 1, "product": 2, "amount": 5})

	fmt.Println(c.Has("product"))
	fmt.Println(c.Has("product", "amount"))
	fmt.Println(c.Has("amount", "price"))
	// Output:
	// true
	// true
	// false
}

func TestHas(t *testing.T) {
	if !collect.Of(1).Keyed().Has() {
		t.Error("Has() of no keys is false")
	}
	if !endless().Keyed().Has(3, 5) {
		t.Error("Has of an endless source is false")
	}
}

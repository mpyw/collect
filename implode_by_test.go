package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of implode.
func ExampleCollection_ImplodeBy() {
	type product struct{ Name string }
	products := collect.Of(product{"Desk"}, product{"Chair"})

	fmt.Println(products.ImplodeBy(func(p product) string { return p.Name }, ", "))
	fmt.Println(collect.Of("1", "2", "3").ImplodeBy(collect.Identity, "-"))
	// Output:
	// Desk, Chair
	// 1-2-3
}

func TestImplodeByOfNothing(t *testing.T) {
	if got := collect.Of[string]().ImplodeBy(collect.Identity, ", "); got != "" {
		t.Errorf("Implode() of nothing = %q", got)
	}
}

func ExampleKeyedCollection_ImplodeBy() {
	c := keyedOf([]string{"a", "b"}, []int{1, 2})

	fmt.Println(c.ImplodeBy(func(k string, v int) string { return fmt.Sprintf("%s=%d", k, v) }, "&"))
	// Output: a=1&b=2
}

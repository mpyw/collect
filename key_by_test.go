package collect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of keyBy.
func ExampleCollection_KeyBy() {
	type product struct {
		ID   string
		Name string
	}
	byID := collect.Of(product{"prod-100", "Desk"}, product{"prod-200", "Chair"}).
		KeyBy(func(p product) string { return p.ID })

	for id, p := range byID.Seq() {
		fmt.Println(id, p.Name)
	}
	// Output:
	// prod-100 Desk
	// prod-200 Chair
}

func TestKeyByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().KeyBy(collect.Identity).Seq())
}

func ExampleKeyedCollection_KeyBy() {
	c := collect.FromMap(map[string]string{"a": "apple"}).
		KeyBy(func(k, v string) string { return strings.ToUpper(k) })

	fmt.Println(c.All())
	// Output: map[A:apple]
}

func TestKeyedKeyByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().KeyBy(func(_, v int) int { return v }).Seq())
}

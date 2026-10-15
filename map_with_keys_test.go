package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of mapWithKeys.
func ExampleCollection_MapWithKeys() {
	type person struct{ Name, Department, Email string }
	byEmail := collect.Of(
		person{"John", "Sales", "john@example.com"},
		person{"Jane", "Marketing", "jane@example.com"},
	).MapWithKeys(func(p person) (string, string) { return p.Email, p.Name })

	fmt.Println(byEmail.All())
	// Output: map[jane@example.com:Jane john@example.com:John]
}

func ExampleKeyedCollection_MapWithKeys() {
	ages := collect.FromMap(map[string]int{"Alice": 31})
	byAge := ages.MapWithKeys(func(name string, age int) (int, string) { return age, name })

	fmt.Println(byAge.All())
	// Output: map[31:Alice]
}

func TestMapWithKeysStopsEarly(t *testing.T) {
	pair := func(v int) (int, int) { return v, v }
	assertStopsEarly2(t, endless().MapWithKeys(pair).Seq())
	assertStopsEarly2(t, endless().Keyed().MapWithKeys(func(k, v int) (int, int) { return v, k }).Seq())
}

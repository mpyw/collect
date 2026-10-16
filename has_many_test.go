package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_HasMany() {
	fmt.Println(collect.Of(1).HasMany(), collect.Of(1, 2).HasMany())
	// Output: false true
}

func ExampleKeyedCollection_HasMany() {
	fmt.Println(collect.FromMap(map[string]int{"a": 1, "b": 2}).HasMany())
	// Output: true
}

func TestHasMany(t *testing.T) {
	if !endless().HasMany() || !endless().Keyed().HasMany() || collect.Of(1).Keyed().HasMany() {
		t.Error("HasMany is wrong")
	}
}

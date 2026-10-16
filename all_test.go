package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_All() {
	fmt.Println(collect.Of(1, 2, 3).All())
	// Output: [1 2 3]
}

func TestAllOfNothing(t *testing.T) {
	assertEqual(t, collect.Of[int]().All(), nil)
}

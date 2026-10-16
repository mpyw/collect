package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

func ExampleCollection_IsNotEmpty() {
	fmt.Println(collect.Of[int]().IsNotEmpty(), collect.Of(1).IsNotEmpty())
	// Output: false true
}

func ExampleKeyedCollection_IsNotEmpty() {
	fmt.Println(collect.FromMap(map[string]int{"a": 1}).IsNotEmpty())
	// Output: true
}

package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

func ExampleCollection_Seq() {
	for v := range collect.Of("a", "b").Seq() {
		fmt.Println(v)
	}
	// Output:
	// a
	// b
}

func ExampleKeyedCollection_Seq() {
	for k, v := range collect.Of("a", "b").Keyed().Seq() {
		fmt.Println(k, v)
	}
	// Output:
	// 0 a
	// 1 b
}

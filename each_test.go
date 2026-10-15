package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

func ExampleCollection_Each() {
	collect.Of("a", "b").Each(func(s string) { fmt.Println(s) })
	// Output:
	// a
	// b
}

func ExampleKeyedCollection_Each() {
	collect.Of("a", "b").Keyed().Each(func(i int, s string) { fmt.Println(i, s) })
	// Output:
	// 0 a
	// 1 b
}

package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

func ExampleCollection_JoinBy() {
	fmt.Println(collect.Of("a", "b", "c").JoinBy(collect.Identity, "/"))
	// Output: a/b/c
}

func ExampleKeyedCollection_JoinBy() {
	c := keyedOf([]string{"a", "b"}, []int{1, 2})

	fmt.Println(c.JoinBy(func(k string, _ int) string { return k }, "/"))
	// Output: a/b
}

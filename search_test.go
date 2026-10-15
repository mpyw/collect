package collect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Search() {
	i, ok := collect.Of(2, 4, 6, 8).Search(func(n int) bool { return n == 4 })

	fmt.Println(i, ok)
	// Output: 1 true
}

func ExampleKeyedCollection_Search() {
	c := collect.Of("Taylor", "Abigail").KeyBy(func(s string) string { return strings.ToLower(s[:1]) })
	k, ok := c.Search(func(_, name string) bool { return name == "Abigail" })

	fmt.Println(k, ok)
	// Output: a true
}

func TestSearchNotFound(t *testing.T) {
	if _, ok := collect.Of(2, 4).Search(func(v int) bool { return v == 5 }); ok {
		t.Error("Search(5) is ok")
	}
}

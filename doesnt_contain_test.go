package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

func ExampleCollection_DoesntContain() {
	c := collect.Of(1, 2, 3, 4, 5)

	fmt.Println(c.DoesntContain(func(n int) bool { return n < 5 }))
	// Output: false
}

func ExampleKeyedCollection_DoesntContain() {
	c := collect.FromMap(map[string]string{"name": "Desk"})

	fmt.Println(c.DoesntContain(func(_, v string) bool { return v == "Table" }))
	// Output: true
}

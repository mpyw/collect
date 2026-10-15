package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// A key that comes more than once is there each time.
func ExampleKeyedCollection_Keys() {
	keys := collect.Of("apple", "avocado", "banana").
		KeyBy(func(s string) string { return s[:1] }).
		Keys().
		All()

	fmt.Println(keys)
	// Output: [a a b]
}

func TestKeysStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Keyed().Keys().Seq())
}

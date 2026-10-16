package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of intersectByKeys.
func ExampleKeyedCollection_IntersectKeys() {
	c := keyedOf([]string{"serial", "type", "year"}, []any{"UX301", "screen", 2009})
	got := c.IntersectKeys(keyedOf([]string{"reference", "type", "year"}, []any{"UX404", "tab", 2011}))

	fmt.Println(got.All())
	// Output: map[type:screen year:2009]
}

func TestIntersectKeysStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().IntersectKeys(collect.Range(1, 9).Keyed()).Seq())
}

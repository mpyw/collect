package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of intersectByKeys.
func ExampleKeyedCollection_IntersectByKeys() {
	c := keyedOf([]string{"serial", "type", "year"}, []any{"UX301", "screen", 2009})
	got := c.IntersectByKeys(keyedOf([]string{"reference", "type", "year"}, []any{"UX404", "tab", 2011}))

	fmt.Println(got.All())
	// Output: map[type:screen year:2009]
}

func TestIntersectByByKeysStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().IntersectByKeys(collect.Range(1, 9).Keyed()).Seq())
}

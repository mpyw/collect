package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of flip.
func ExampleKeyedCollection_FlipBy() {
	c := collect.FromMap(map[string]string{"name": "taylor", "framework": "laravel"})

	fmt.Println(c.FlipBy(collect.Identity).ToMap())
	// Output: map[laravel:framework taylor:name]
}

func TestFlipByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().FlipBy(collect.Identity).Seq())
}

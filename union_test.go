package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of union.
func ExampleKeyedCollection_Union() {
	c := keyedOf([]int{1, 2}, []string{"a", "b"}).
		Union(keyedOf([]int{3, 1}, []string{"c", "b"}))

	fmt.Println(c.All())
	// Output: map[1:a 2:b 3:c]
}

func TestUnionKeepsOrder(t *testing.T) {
	got := keyedOf([]int{1, 2}, []string{"a", "b"}).Union(keyedOf([]int{3, 1, 3}, []string{"c", "b", "d"}))
	assertPairs(t, got, []int{1, 2, 3}, []string{"a", "b", "c"})
}

func TestUnionStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().Union(collect.Of(1).Keyed()).Seq())
	assertStopsEarly2(t, collect.Of(1).Keyed().Union(endless().Keyed()).Seq())
}

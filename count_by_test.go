package collect_test

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of countBy.
func ExampleCollection_CountBy() {
	fmt.Println(collect.Of(1, 2, 2, 2, 3).CountBy(collect.Identity).ToMap())

	domain := func(email string) string { return email[strings.Index(email, "@")+1:] }
	fmt.Println(collect.Of("alice@gmail.com", "bob@yahoo.com", "carlos@gmail.com").CountBy(domain).ToMap())
	// Output:
	// map[1:1 2:3 3:1]
	// map[gmail.com:2 yahoo.com:1]
}

func TestCountByKeepsOrder(t *testing.T) {
	assertPairs(t, collect.Of("b", "a", "b").CountBy(collect.Identity), []string{"b", "a"}, []int{2, 1})
}

func TestCountByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Range(1, 9).CountBy(collect.Identity).Seq())
}

func ExampleKeyedCollection_CountBy() {
	c := collect.FromSeq2(slices.All([]string{"apple", "avocado", "banana"}))

	fmt.Println(c.CountBy(func(_ int, s string) string { return s[:1] }).ToMap())
	// Output: map[a:2 b:1]
}

func TestKeyedCountByStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.Range(1, 9).Keyed().CountBy(func(k, _ int) int { return k }).Seq())
}

func TestCountByNaN(t *testing.T) {
	nan := math.NaN()
	assertEqual(t, collect.Of(nan, nan, 1).CountBy(collect.Identity).Values().All(), []int{1, 1, 1})
	assertEqual(t, collect.Of(nan, 1).Keyed().CountBy(collect.Value).Values().All(), []int{1, 1})
}

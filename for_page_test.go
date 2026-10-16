package collect_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of forPage.
func ExampleCollection_ForPage() {
	fmt.Println(collect.Range(1, 9).ForPage(2, 3).All())
	// Output: [4 5 6]
}

func TestForPage(t *testing.T) {
	c := collect.Range(1, 9)
	assertEqual(t, c.ForPage(0, 3).All(), []int{1, 2, 3})
	assertEqual(t, c.ForPage(4, 3).All(), nil)
}

func ExampleKeyedCollection_ForPage() {
	fmt.Println(collect.Range(1, 9).Keyed().ForPage(2, 3).Keys().All())
	// Output: [3 4 5]
}

// A negative perPage counts from the end, as Laravel's forPage does, and a
// page far past the end is empty.
func TestForPageEdges(t *testing.T) {
	c := collect.Range(1, 7)
	assertEqual(t, c.ForPage(2, -3).All(), []int{1, 2, 3, 4})
	assertEqual(t, c.ForPage(0, -3).All(), []int{1, 2, 3, 4})
	assertEqual(t, c.ForPage(math.MaxInt, 2).All(), nil)
	assertEqual(t, c.Keyed().ForPage(math.MaxInt, 2).Keys().All(), nil)
}

func TestForPageStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().ForPage(2, 10).Seq())
	assertStopsEarly2(t, endless().Keyed().ForPage(2, 10).Seq())
}

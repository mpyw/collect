package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of crossJoin.
func ExampleCollection_CrossJoin() {
	for p := range collect.Of(1, 2).CrossJoin(collect.Of("a", "b")).Seq() {
		fmt.Println(p.A, p.B)
	}
	// Output:
	// 1 a
	// 1 b
	// 2 a
	// 2 b
}

func TestCrossJoinStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().CrossJoin(collect.Of(1, 2)).Seq())
}

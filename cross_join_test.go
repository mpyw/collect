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

// An empty other gives no values, and reads none of an endless collection.
func TestCrossJoinEmptyOther(t *testing.T) {
	seq, n := counted()
	assertEqual(t, collect.FromSeq(seq).CrossJoin(collect.Of[string]()).All(), nil)
	if *n != 0 {
		t.Errorf("read %d values, want 0", *n)
	}
}

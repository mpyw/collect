package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Eager reads a source that can be read only once, and the result can be
// read any number of times.
func ExampleCollection_Eager() {
	c := once(1, 2).Eager()

	fmt.Println(c.All(), c.All())
	// Output: [1 2] [1 2]
}

func ExampleKeyedCollection_Eager() {
	c := once("a", "b").Keyed().Eager()

	fmt.Println(c.Keys().All(), c.Keys().All())
	// Output: [0 1] [0 1]
}

func TestEagerKeepsEveryKey(t *testing.T) {
	k := once(1, 2).KeyBy(func(int) string { return "same" }).Eager()
	assertPairs(t, k, []string{"same", "same"}, []int{1, 2})
	assertPairs(t, k.Take(1), []string{"same"}, []int{1})
}

// Eager reads when it is called.
func TestEagerReadsNow(t *testing.T) {
	seq, n := counted()
	collect.FromSeq(seq).Take(3).Eager()
	if *n != 3 {
		t.Errorf("Eager read %d values, want 3", *n)
	}
}

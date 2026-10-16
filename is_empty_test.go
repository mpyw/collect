package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_IsEmpty() {
	fmt.Println(collect.Of[int]().IsEmpty(), collect.Of(1).IsEmpty())
	// Output: true false
}

func ExampleKeyedCollection_IsEmpty() {
	fmt.Println(collect.FromMap(map[string]int{}).IsEmpty())
	// Output: true
}

func TestIsEmptyReadsOne(t *testing.T) {
	seq, n := counted()
	collect.FromSeq(seq).IsEmpty()
	if *n != 1 {
		t.Errorf("IsEmpty read %d values, want 1", *n)
	}
	if collect.Of(1).Keyed().IsEmpty() {
		t.Error("IsEmpty of a pair is true")
	}
}

// A source that can be read once loses the value IsEmpty reads, unless
// Remember keeps it.
func ExampleCollection_IsEmpty_remember() {
	cursor := func() func(func(int) bool) {
		used := false
		return func(yield func(int) bool) {
			if used {
				return
			}
			used = true
			for n := 1; n <= 3; n++ {
				if !yield(n) {
					return
				}
			}
		}
	}

	once := collect.FromSeq(cursor())
	fmt.Println(once.IsEmpty(), once.All()) // false [], the cursor is used up

	kept := collect.FromSeq(cursor()).Remember()
	fmt.Println(kept.IsEmpty(), kept.All()) // false [1 2 3]

	// Output:
	// false []
	// false [1 2 3]
}

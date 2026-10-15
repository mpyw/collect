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

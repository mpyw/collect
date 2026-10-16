package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_HasSole() {
	fmt.Println(collect.Of(1).HasSole(), collect.Of(1, 2).HasSole())
	// Filter first to ask about the values that match.
	fmt.Println(collect.Of(1, 2).Filter(func(n int) bool { return n > 1 }).HasSole())
	// Output:
	// true false
	// true
}

func ExampleKeyedCollection_HasSole() {
	fmt.Println(collect.FromMap(map[string]int{"a": 1}).HasSole())
	// Output: true
}

func TestHasSole(t *testing.T) {
	if collect.Of[int]().HasSole() || collect.Of[int]().Keyed().HasSole() || collect.Of(1, 2).Keyed().HasSole() {
		t.Error("HasSole of none or two is true")
	}
	seq, n := counted()
	collect.FromSeq(seq).HasSole()
	if *n != 2 {
		t.Errorf("HasSole read %d values, want 2", *n)
	}
}

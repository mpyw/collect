package collect_test

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_RandomN() {
	c, err := collect.Range(1, 5).RandomN(3)
	fmt.Println(c.Count(), err)

	_, err = collect.Range(1, 5).RandomN(6)
	fmt.Println(err)
	// Output:
	// 3 <nil>
	// collect: not enough items: asked for 6, but there are only 5
}

func TestRandomN(t *testing.T) {
	c, err := collect.Range(1, 9).RandomN(4)
	if err != nil {
		t.Fatal(err)
	}
	got := c.All()
	if len(got) != 4 || !slices.IsSorted(got) || collect.FromSlice(got).UniqueBy(collect.Identity).Count() != 4 {
		t.Errorf("RandomN(4) = %v, want 4 distinct values in their order", got)
	}
	if c, err := collect.Of(1).RandomN(0); err != nil || c.Count() != 0 {
		t.Errorf("RandomN(0) = %v, %v", c.All(), err)
	}
	if _, err := collect.Of(1).RandomN(2); !errors.Is(err, collect.ErrNotEnoughItems) {
		t.Errorf("RandomN(2) of 1: %v", err)
	}
	// A negative count is 0, as in Laravel.
	if c, err := collect.Of(1).RandomN(-1); err != nil || c.Count() != 0 {
		t.Errorf("RandomN(-1) = %v, %v", c.All(), err)
	}
}

func ExampleKeyedCollection_RandomN() {
	c, err := collect.Range(1, 5).Keyed().RandomN(3)

	fmt.Println(c.Count(), err)
	// Output: 3 <nil>
}

func TestKeyedRandomN(t *testing.T) {
	c, err := collect.Of(10, 20, 30, 40).Keyed().RandomN(2)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range c.Seq() {
		if v != (k+1)*10 {
			t.Errorf("RandomN() lost the key of %d: %d", v, k)
		}
	}
	if !slices.IsSorted(c.Keys().All()) {
		t.Errorf("RandomN() = %v, want the order they come", c.Keys().All())
	}
	if _, err := collect.Of(1).Keyed().RandomN(2); !errors.Is(err, collect.ErrNotEnoughItems) {
		t.Errorf("RandomN(2) of 1: %v", err)
	}
	if c, err := collect.Of(1).Keyed().RandomN(-1); err != nil || c.Count() != 0 {
		t.Errorf("RandomN(-1) = %v, %v", c.Keys().All(), err)
	}
}

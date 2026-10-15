package collect_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Sole() {
	v, err := collect.Of(1, 2, 3).Filter(func(n int) bool { return n == 2 }).Sole()
	fmt.Println(v, err)

	_, err = collect.Of(1, 2, 3).Sole()
	fmt.Println(errors.Is(err, collect.ErrMultipleItemsFound))
	// Output:
	// 2 <nil>
	// true
}

func TestSole(t *testing.T) {
	if _, err := collect.Of[int]().Sole(); !errors.Is(err, collect.ErrItemNotFound) {
		t.Errorf("Sole() of nothing: %v", err)
	}
	if _, err := endless().Sole(); !errors.Is(err, collect.ErrMultipleItemsFound) {
		t.Errorf("Sole() of an endless source: %v", err)
	}
}

func ExampleKeyedCollection_Sole() {
	k, v, err := collect.FromMap(map[string]int{"only": 1}).Sole()

	fmt.Println(k, v, err)
	// Output: only 1 <nil>
}

func TestKeyedSole(t *testing.T) {
	if _, _, err := collect.Of[int]().Keyed().Sole(); !errors.Is(err, collect.ErrItemNotFound) {
		t.Errorf("Sole() of nothing: %v", err)
	}
	if _, _, err := endless().Keyed().Sole(); !errors.Is(err, collect.ErrMultipleItemsFound) {
		t.Errorf("Sole() of many: %v", err)
	}
}

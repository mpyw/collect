package collect_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Ensure() {
	ints, err := collect.Of[any](1, 2).Ensure[int]()
	fmt.Println(ints.All(), err)

	_, err = collect.Of[any](1, "a").Ensure[int]()
	fmt.Println(err)
	// Output:
	// [1 2] <nil>
	// collect: unexpected value: a value of type string is not int
}

func TestEnsureError(t *testing.T) {
	_, err := collect.Of[any]("a").Ensure[int]()
	if !errors.Is(err, collect.ErrUnexpectedValue) {
		t.Errorf("Ensure() error = %v", err)
	}
}

func ExampleKeyedCollection_Ensure() {
	c, err := collect.FromMap(map[string]any{"a": 1}).Ensure[int]()

	fmt.Println(c.ToMap(), err)
	// Output: map[a:1] <nil>
}

func TestKeyedEnsureError(t *testing.T) {
	_, err := collect.FromMap(map[string]any{"a": "x"}).Ensure[int]()
	if !errors.Is(err, collect.ErrUnexpectedValue) {
		t.Errorf("Ensure() error = %v", err)
	}
}

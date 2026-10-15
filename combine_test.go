package collect_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of combine.
func ExampleCombine() {
	c, err := collect.Combine(collect.Of("name", "age"), collect.Of[any]("George", 29))
	fmt.Println(c.All(), err)

	_, err = collect.Combine(collect.Of("name", "age"), collect.Of[any]("George"))
	fmt.Println(err)
	// Output:
	// map[age:29 name:George] <nil>
	// collect: lengths do not match: 2 keys and 1 values
}

func TestCombineError(t *testing.T) {
	if _, err := collect.Combine(collect.Of("a"), collect.Of(1, 2)); !errors.Is(err, collect.ErrLengthMismatch) {
		t.Errorf("Combine() of 1 key and 2 values: %v", err)
	}
}

func TestCombineStopsEarly(t *testing.T) {
	c, err := collect.Combine(collect.Range(1, 9), collect.Range(1, 9))
	if err != nil {
		t.Fatal(err)
	}
	assertStopsEarly2(t, c.Seq())
}

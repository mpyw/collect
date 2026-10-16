package collect_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_MustEnsure() {
	defer func() { fmt.Println("panic:", recover()) }()

	for n := range collect.Of[any](1, "a").MustEnsure[int]().Seq() {
		fmt.Println(n)
	}
	// Output:
	// 1
	// panic: collect: unexpected value: a value of type string is not int
}

func TestMustEnsure(t *testing.T) {
	assertEqual(t, collect.Of[any](1, 2).MustEnsure[int]().All(), []int{1, 2})
	c := collect.Of[any](1, "a").MustEnsure[int]()
	defer func() {
		err, ok := recover().(error)
		if !ok || !errors.Is(err, collect.ErrUnexpectedValue) {
			t.Errorf("MustEnsure panicked with %v", err)
		}
	}()
	c.All()
}

func TestMustEnsureStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().MustEnsure[int]().Seq())
}

func ExampleKeyedCollection_MustEnsure() {
	fmt.Println(collect.FromMap(map[string]any{"a": 1}).MustEnsure[int]().ToMap())
	// Output: map[a:1]
}

func TestKeyedMustEnsure(t *testing.T) {
	assertPanicsWith(t, collect.ErrUnexpectedValue, func() {
		collect.Of[any]("x").Keyed().MustEnsure[int]().ToMap()
	})
	assertStopsEarly2(t, endless().Keyed().MustEnsure[int]().Seq())
}

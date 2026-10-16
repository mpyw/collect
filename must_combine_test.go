package collect_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleMustCombine() {
	fmt.Println(collect.MustCombine(collect.Of("a", "b"), collect.Range(1, 2)).ToMap())
	// Output: map[a:1 b:2]
}

func TestMustCombine(t *testing.T) {
	assertPairs(t, collect.MustCombine(collect.Of("a", "b"), collect.Of(1, 2)), []string{"a", "b"}, []int{1, 2})
	for _, values := range []collect.Collection[int]{collect.Of(1), collect.Of(1, 2, 3)} {
		func() {
			defer func() {
				err, ok := recover().(error)
				if !ok || !errors.Is(err, collect.ErrLengthMismatch) {
					t.Errorf("MustCombine of %d values panicked with %v", values.Count(), err)
				}
			}()
			collect.MustCombine(collect.Of("a", "b"), values).ToMap()
		}()
	}
}

// MustCombine is lazy: an endless source is fine until a read stops.
func TestMustCombineStopsEarly(t *testing.T) {
	assertStopsEarly2(t, collect.MustCombine(endless(), endless()).Seq())
}

package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Keyed() {
	for k, v := range collect.Of("a", "b", "c").Keyed().Seq() {
		fmt.Println(k, v)
	}
	// Output:
	// 0 a
	// 1 b
	// 2 c
}

func TestKeyedStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().Seq())
}

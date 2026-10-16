package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExamplePairs() {
	c := collect.Of("apple", "avocado", "banana").KeyBy(func(s string) string { return s[:1] })

	for p := range collect.Pairs(c).Seq() {
		fmt.Println(p.First, p.Second)
	}
	// Output:
	// a apple
	// a avocado
	// b banana
}

func TestPairsStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Pairs(endless().Keyed()).Seq())
}

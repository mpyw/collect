package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_Random() {
	n, ok := collect.Range(1, 5).Random()

	fmt.Println(n >= 1 && n <= 5, ok)
	// Output: true true
}

func TestRandom(t *testing.T) {
	seen := map[int]bool{}
	for range 200 {
		v, ok := collect.Range(1, 3).Random()
		if !ok || v < 1 || v > 3 {
			t.Fatalf("Random() = %v, %v", v, ok)
		}
		seen[v] = true
	}
	if len(seen) != 3 {
		t.Errorf("200 draws chose only %v", seen)
	}
	if _, ok := collect.Of[int]().Random(); ok {
		t.Error("Random() of nothing is ok")
	}
}

func ExampleKeyedCollection_Random() {
	k, v, ok := collect.Of("a", "b").Keyed().Random()

	fmt.Println((k == 0 && v == "a") || (k == 1 && v == "b"), ok)
	// Output: true true
}

func TestKeyedRandomOfNothing(t *testing.T) {
	if _, _, ok := collect.Of[int]().Keyed().Random(); ok {
		t.Error("Random() of nothing is ok")
	}
}

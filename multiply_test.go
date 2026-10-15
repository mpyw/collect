package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of multiply.
func ExampleCollection_Multiply() {
	fmt.Println(collect.Of("a", "b").Multiply(3).All())
	// Output: [a b a b a b]
}

func TestMultiplyNone(t *testing.T) {
	assertEqual(t, collect.Of(1, 2).Multiply(0).All(), nil)
	assertEqual(t, collect.Of(1, 2).Multiply(-1).All(), nil)
}

func TestMultiplyStopsEarly(t *testing.T) {
	assertStopsEarly(t, collect.Of(1, 2).Multiply(5).Seq())
	assertStopsEarly(t, endless().Multiply(2).Seq())
}

func ExampleKeyedCollection_Multiply() {
	fmt.Println(collect.Of("a", "b").Keyed().Multiply(2).Keys().All())
	// Output: [0 1 0 1]
}

func TestKeyedMultiply(t *testing.T) {
	assertEqual(t, collect.Of(1).Keyed().Multiply(-1).Keys().All(), nil)
	assertStopsEarly2(t, collect.Of(1, 2).Keyed().Multiply(5).Seq())
	assertStopsEarly2(t, endless().Keyed().Multiply(2).Seq())
}

// Multiply reads the collection once, so a source that can be read once
// repeats.
func TestMultiplyReadsOnce(t *testing.T) {
	assertEqual(t, once(1, 2).Multiply(3).All(), []int{1, 2, 1, 2, 1, 2})
	assertEqual(t, once("a").Keyed().Multiply(2).Values().All(), []string{"a", "a"})
	seq, n := counted()
	collect.FromSeq(seq).Take(2).Multiply(3).All()
	if *n != 2 {
		t.Errorf("Multiply(3) read %d values of 2, want 2", *n)
	}
}

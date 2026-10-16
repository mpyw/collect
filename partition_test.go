package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of partition.
func ExampleCollection_Partition() {
	under3, over3 := collect.Of(1, 2, 3, 4, 5, 6).Partition(func(n int) bool { return n < 3 })

	fmt.Println(under3.All(), over3.All())
	// Output: [1 2] [3 4 5 6]
}

func ExampleKeyedCollection_Partition() {
	even, odd := collect.Of("a", "b", "c").Keyed().Partition(func(i int, _ string) bool { return i%2 == 0 })

	fmt.Println(even.Keys().All(), odd.Keys().All())
	// Output: [0 2] [1]
}

// Partition reads the collection once, so a source that can be read once
// fills both halves.
func TestPartitionReadsOnce(t *testing.T) {
	small, big := once(1, 2, 3, 4).Partition(func(v int) bool { return v < 3 })
	assertEqual(t, big.All(), []int{3, 4})
	assertEqual(t, small.All(), []int{1, 2})
	assertEqual(t, small.All(), []int{1, 2})
	even, odd := once("a", "b", "c").Keyed().Partition(func(k int, _ string) bool { return k%2 == 0 })
	assertPairs(t, odd, []int{1}, []string{"b"})
	assertPairs(t, even, []int{0, 2}, []string{"a", "c"})
}

func TestPartitionStopsEarly(t *testing.T) {
	in, out := collect.Range(1, 9).Partition(func(v int) bool { return v%2 == 0 })
	assertStopsEarly(t, out.Seq())
	assertStopsEarly(t, in.Seq())
	kin, kout := collect.Range(1, 9).Keyed().Partition(func(k, _ int) bool { return k%2 == 0 })
	assertStopsEarly2(t, kin.Seq())
	assertStopsEarly2(t, kout.Seq())
}

// A panic keeps nothing, so the next read starts again, rather than return
// what was read before the panic.
func TestPartitionAfterPanic(t *testing.T) {
	calls := 0
	in, out := collect.Of(1, 2, 3, 4).Partition(func(v int) bool {
		if calls++; calls == 3 {
			panic("boom")
		}
		return v%2 == 0
	})
	func() {
		defer func() { _ = recover() }()
		in.All()
	}()
	assertEqual(t, in.All(), []int{2, 4})
	assertEqual(t, out.All(), []int{1, 3})
	kin, _ := collect.Of(1, 2).Keyed().Partition(func(k, _ int) bool {
		if calls++; calls == 9 {
			panic("boom")
		}
		return true
	})
	func() {
		defer func() { _ = recover() }()
		kin.All()
	}()
	assertPairs(t, kin, []int{0, 1}, []int{1, 2})
}

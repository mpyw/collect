package collect_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/mpyw/collect"
)

// Remember reads the source once. A later read gets the kept values, then
// reads on.
func ExampleCollection_Remember() {
	c := collect.Range(1, 5).
		TapEach(func(n int) { fmt.Println("read", n) }).
		Remember()

	fmt.Println(c.Take(2).All())
	fmt.Println(c.Take(3).All())
	// Output:
	// read 1
	// read 2
	// [1 2]
	// read 3
	// [1 2 3]
}

func ExampleKeyedCollection_Remember() {
	c := once("a", "b").Keyed().Remember()

	fmt.Println(c.Keys().All(), c.Keys().All())
	// Output: [0 1] [0 1]
}

func TestRememberAfterPart(t *testing.T) {
	k := once("a", "b").Keyed().Remember()
	assertPairs(t, k.Take(1), []int{0}, []string{"a"})
	assertPairs(t, k, []int{0, 1}, []string{"a", "b"})
}

// Remember reads each value of the source once, however many reads ask.
func TestRememberReadsOnce(t *testing.T) {
	seq, n := counted()
	c := collect.FromSeq(seq).Remember()
	assertEqual(t, c.Take(2).All(), []int{1, 2})
	assertEqual(t, c.Take(3).All(), []int{1, 2, 3})
	if *n != 3 {
		t.Errorf("read %d values of the source, want 3", *n)
	}
}

func TestRememberStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().Remember().Seq())
	assertStopsEarly2(t, endless().Keyed().Remember().Seq())
}

// A source that is not read to its end is stopped once the collection is
// garbage collected.
func TestRememberStopsSourceWhenCollected(t *testing.T) {
	for name, remember := range map[string]func(func()){
		"Collection": func(stopped func()) {
			c := collect.FromSeq(func(yield func(int) bool) {
				defer stopped()
				for i := 0; yield(i); i++ {
				}
			}).Remember()
			c.Take(1).All()
		},
		"KeyedCollection": func(stopped func()) {
			c := collect.FromSeq2(func(yield func(int, int) bool) {
				defer stopped()
				for i := 0; yield(i, i); i++ {
				}
			}).Remember()
			c.Take(1).Keys().All()
		},
	} {
		t.Run(name, func(t *testing.T) {
			done := make(chan struct{})
			remember(func() { close(done) })
			for range 100 {
				runtime.GC()
				select {
				case <-done:
					return
				case <-time.After(10 * time.Millisecond):
				}
			}
			t.Error("the source was not stopped")
		})
	}
}

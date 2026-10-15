package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of join, with a final glue.
func ExampleCollection_JoinFinalBy() {
	fmt.Println(collect.Of("a", "b", "c").JoinFinalBy(collect.Identity, ", ", " and "))
	fmt.Println(collect.Of("a", "b").JoinFinalBy(collect.Identity, ", ", " and "))
	fmt.Println(collect.Of("a").JoinFinalBy(collect.Identity, ", ", " and "))
	// Output:
	// a, b and c
	// a and b
	// a
}

func TestJoinByFinal(t *testing.T) {
	if got := collect.Of[string]().JoinFinalBy(collect.Identity, ", ", " and "); got != "" {
		t.Errorf("JoinFinal() of nothing = %q", got)
	}
	if got := collect.Of("a", "b").JoinFinalBy(collect.Identity, ", ", ""); got != "a, b" {
		t.Errorf("JoinFinal() with no final glue = %q", got)
	}
}

func ExampleKeyedCollection_JoinFinalBy() {
	c := keyedOf([]string{"a", "b", "c"}, []int{1, 2, 3})

	fmt.Println(c.JoinFinalBy(func(k string, _ int) string { return k }, ", ", " and "))
	// Output: a, b and c
}

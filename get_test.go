package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's example of get.
func ExampleKeyedCollection_Get() {
	c := collect.FromMap(map[string]string{"name": "taylor", "framework": "laravel"})

	fmt.Println(c.Get("name"))
	fmt.Println(c.Get("age"))
	// Output:
	// taylor true
	//  false
}

// Get returns the first value of a key, and All keeps the last.
func TestGetFirstOfAKey(t *testing.T) {
	k := keyedOf([]string{"a", "a"}, []int{1, 2})
	if v, _ := k.Get("a"); v != 1 {
		t.Errorf(`Get("a") = %v, want the first value`, v)
	}
	if v, ok := endless().Keyed().Get(5); v != 6 || !ok {
		t.Errorf("Get(5) of an endless source = %v, %v", v, ok)
	}
}

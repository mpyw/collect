package collect_test

import (
	"fmt"

	"github.com/mpyw/collect"
)

// Laravel's example of reduce.
func ExampleCollection_Reduce() {
	total := collect.Of(1, 2, 3).Reduce(func(carry, n int) int { return carry + n }, 4)

	fmt.Println(total)
	// Output: 10
}

// Laravel's example of reduceWithKeys.
func ExampleKeyedCollection_Reduce() {
	prices := collect.MustCombine(collect.Of("usd", "gbp", "eur"), collect.Of(1400, 1200, 1000))
	rates := map[string]int{"usd": 1, "gbp": 1, "eur": 1}
	total := prices.Reduce(func(carry int, currency string, n int) int { return carry + n*rates[currency] }, 0)

	fmt.Println(total)
	// Output: 3600
}

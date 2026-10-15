package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

// Laravel's examples of sum.
func ExampleCollection_SumBy() {
	type book struct {
		Name  string
		Pages int
	}
	books := collect.Of(book{"JavaScript: The Good Parts", 176}, book{"JavaScript: The Definitive Guide", 1096})

	fmt.Println(collect.Of(1, 2, 3, 4, 5).SumBy(collect.Identity))
	fmt.Println(books.SumBy(func(b book) int { return b.Pages }))
	// Output:
	// 15
	// 1272
}

func TestSumByOfNothing(t *testing.T) {
	if got := collect.Of[float64]().SumBy(collect.Identity); got != 0 {
		t.Errorf("Sum() of nothing = %v", got)
	}
}

func ExampleKeyedCollection_SumBy() {
	prices := collect.FromMap(map[string]int{"desk": 200, "chair": 100})

	fmt.Println(prices.SumBy(func(_ string, price int) int { return price }))
	// Output: 300
}

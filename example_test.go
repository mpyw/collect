package collect_test

import (
	"errors"
	"fmt"

	"github.com/mpyw/collect"
)

func Example() {
	type order struct {
		Customer string
		Item     string
		Price    int
	}
	orders := collect.Of(
		order{"alice", "desk", 200},
		order{"bob", "chair", 100},
		order{"alice", "lamp", 40},
		order{"carol", "sofa", 300},
		order{"bob", "pen", 5},
	)
	price := func(o order) int { return o.Price }
	customer := func(o order) string { return o.Customer }
	item := func(o order) string { return o.Item }

	// Chain Filter and Map. Map may change the type of the values.
	cheap := orders.
		Filter(func(o order) bool { return o.Price < 100 }).
		Map(func(o order) string { return o.Item }).
		All()
	fmt.Println("cheap:", cheap) // cheap: [lamp pen]

	// A method that needs numbers or an order takes a key function.
	fmt.Println("total:", orders.SumBy(price)) // total: 645
	top, ok := orders.SortByDesc(price).First()
	fmt.Println("top:", top.Item, ok) // top: sofa true

	// GroupBy nests: each group is a Collection with every method.
	for name, group := range orders.GroupBy(customer).Seq() {
		fmt.Println(name, group.Count(), group.SumBy(price)) // alice 2 240, then bob 2 105, then carol 1 300
	}

	// Keys are opt-in. KeyBy gives each value one.
	sofa, _ := orders.KeyBy(item).Get("sofa")
	fmt.Println("sofa:", sofa.Customer) // sofa: carol

	// Identity stands in for the key function when the values are their own keys.
	unique := collect.Of(3, 1, 2, 3, 1).UniqueBy(collect.Identity).SortBy(collect.Identity)
	fmt.Println("unique:", unique.All()) // unique: [1 2 3]

	// Chunk returns a Collection of collections.
	for chunk := range unique.Chunk(2).Seq() {
		fmt.Println("chunk:", chunk.All()) // chunk: [1 2], then chunk: [3]
	}

	// A lookup that may find nothing returns a bool. One that can fail in
	// more ways returns an error.
	_, err := orders.Filter(func(o order) bool { return o.Customer == "dave" }).Sole()
	fmt.Println("dave:", errors.Is(err, collect.ErrItemNotFound)) // dave: true

	// It is lazy, so an endless source is fine while something ends it.
	naturals := func(yield func(int) bool) {
		for n := 1; ; n++ {
			if !yield(n) {
				return
			}
		}
	}
	squares := collect.FromSeq(naturals).Map(func(n int) int { return n * n }).Take(5).All()
	fmt.Println("squares:", squares) // squares: [1 4 9 16 25]

	// Output:
	// cheap: [lamp pen]
	// total: 645
	// top: sofa true
	// alice 2 240
	// bob 2 105
	// carol 1 300
	// sofa: carol
	// unique: [1 2 3]
	// chunk: [1 2]
	// chunk: [3]
	// dave: true
	// squares: [1 4 9 16 25]
}

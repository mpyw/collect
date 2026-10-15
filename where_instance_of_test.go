package collect_test

import (
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_WhereInstanceOf() {
	values := collect.Of[any](1, "a", 2, square(3))

	fmt.Println(values.WhereInstanceOf[int]().All())
	fmt.Println(values.WhereInstanceOf[shape]().Map(shape.area).All())
	// Output:
	// [1 2]
	// [9]
}

type shape interface{ area() int }

type square int

func (s square) area() int { return int(s * s) }

func TestWhereInstanceOf(t *testing.T) {
	values := collect.Of[any](1, "a", 2, square(3))
	assertEqual(t, values.WhereInstanceOf[int]().All(), []int{1, 2})
	shapes := values.WhereInstanceOf[shape]().Map(shape.area).All()
	assertEqual(t, shapes, []int{9})
	assertEqual(t, collect.Of(fmt.Errorf("x")).WhereInstanceOf[fmt.Stringer]().All(), nil)
}

func TestWhereInstanceOfStopsEarly(t *testing.T) {
	assertStopsEarly(t, endless().WhereInstanceOf[int]().Seq())
}

func ExampleKeyedCollection_WhereInstanceOf() {
	c := collect.Of[any](1, "a", 2).Keyed().WhereInstanceOf[int]()

	fmt.Println(c.Keys().All(), c.Values().All())
	// Output: [0 2] [1 2]
}

func TestKeyedWhereInstanceOfStopsEarly(t *testing.T) {
	assertStopsEarly2(t, endless().Keyed().WhereInstanceOf[int]().Seq())
}

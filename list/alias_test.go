package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gvaligiani/al.go/list"
)

// type MyType []int64

// func UpdateValues[T ~int64, L ~[]T](values L) L {
// 	fmt.Printf("before: %v \n", values)
// 	values[0] = values[len(values)-1]
// 	values[len(values)-1] = 0
// 	fmt.Printf("middle: %v \n", values)
// 	values = (values)[:len(values)-1]
// 	fmt.Printf("middle: %v \n", values)
// 	values = append(values, 99)
// 	fmt.Printf("after: %v \n", values)
// 	return values
// }

// func TestXxx(t *testing.T) {
// 	values := MyType{10, 20, 30, 40}
// 	fmt.Printf("before: %v \n", values)
// 	values = UpdateValues(values)
// 	fmt.Printf("after: %v \n", values)
// 	require.True(t, false, "wrong")
// }

func TestAlias(t *testing.T) {

	// builder

	l := list.New(
		Item{Value: 10},
		Item{Value: 12},
		Item{Value: 10}, // duplicate
	)

	// add

	require.True(t, l.Add(Item{Value: 17}), "add 17")
	require.True(t, l.Add(Item{Value: 17}), "add 17 twice")

	// remove

	require.True(t, l.Remove(Item{Value: 10}), "remove 10")
	require.False(t, l.Remove(Item{Value: 10}), "remove 10 twice")

	// list = 12, 17, 17

	// predicate

	require.True(t, l.AllOf(func(i Item) bool { return i.Value < 20 }), "all_of")
	require.False(t, l.AllOf(func(i Item) bool { return i.Value < 10 }), "all_of")
	require.True(t, l.NoneOf(func(i Item) bool { return i.Value < 10 }), "none_of")
	require.False(t, l.AnyOf(func(i Item) bool { return i.Value < 10 }), "any_of")

	// find

	require.True(t, l.Find(Item{Value: 17}), "find 17")
	require.False(t, l.Find(Item{Value: 15}), "find 15")

	item, found := l.FindIf(func(i Item) bool { return i.Value%2 == 0 })
	require.Equal(t, Item{Value: 12}, item, "odd item")
	require.True(t, found, "found odd")

	item, found = l.FindIfNot(func(i Item) bool { return i.Value%2 == 0 })
	require.Equal(t, Item{Value: 17}, item, "even item")
	require.True(t, found, "found even")

	// range

	var sum int64

	sum = 0
	l.Each(func(i Item) { sum += i.Value })
	require.Equal(t, int64(12+17+17), sum, "sum")

	sum = 0
	for _, i := range l {
		sum += i.Value
	}
	require.Equal(t, int64(12+17+17), sum, "sum2")

	// remove if

	odds := l.Copy()
	odds.RemoveIf(func(item Item) bool { return item.Value%2 == 1 })
	assertEquals(t, list.New(Item{Value: 12}), odds, "wrong odds")

	// keep if

	evens := l.Copy()
	evens.KeepIf(func(item Item) bool { return item.Value%2 == 1 })
	assertEquals(t, list.New(Item{Value: 17}, Item{Value: 17}), evens, "wrong evens")

	// check source of copy

	assertEquals(t, list.New(Item{Value: 12}, Item{Value: 17}, Item{Value: 17}), l, "source of copy has been modified")

	// clear

	require.Equal(t, 3, len(l), "len before clear")
	require.False(t, l.IsEmpty(), "is_empty before clear")
	l.Clear()
	require.Equal(t, 0, len(l), "len after clear")
	require.True(t, l.IsEmpty(), "is_empty after clear")
}

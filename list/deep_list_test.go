package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gvaligiani/al.go/list"
)

func TestDeepList(t *testing.T) {

	// builder

	l := list.NewDeep(
		&Item{Value: 10},
		&Item{Value: 12},
		&Item{Value: 10}, // duplicate
	)

	// add

	require.True(t, l.Add(&Item{Value: 17}), "add 17")
	require.True(t, l.Add(&Item{Value: 17}), "add 17 twice")

	// remove

	require.True(t, l.Remove(&Item{Value: 10}), "remove 10")
	require.False(t, l.Remove(&Item{Value: 10}), "remove 10 twice")

	// list = 12, 17, 17

	// predicate

	require.True(t, l.AllOf(func(_ int, i *Item) bool { return i.Value < 20 }), "all_of")
	require.False(t, l.AllOf(func(_ int, i *Item) bool { return i.Value < 10 }), "all_of")
	require.True(t, l.NoneOf(func(_ int, i *Item) bool { return i.Value < 10 }), "none_of")
	require.False(t, l.AnyOf(func(_ int, i *Item) bool { return i.Value < 10 }), "any_of")

	// find

	_, found := l.Find(&Item{Value: 17})
	require.True(t, found, "find 17")
	_, found = l.Find(&Item{Value: 15})
	require.False(t, found, "find 15")

	_, item, found := l.FindIf(func(_ int, i *Item) bool { return i.Value%2 == 0 })
	require.Equal(t, &Item{Value: 12}, item, "odd item")
	require.True(t, found, "found odd")

	_, item, found = l.FindIfNot(func(_ int, i *Item) bool { return i.Value%2 == 0 })
	require.Equal(t, &Item{Value: 17}, item, "even item")
	require.True(t, found, "found even")

	// range

	var sum int64

	sum = 0
	l.Each(func(_ int, i *Item) { sum += i.Value })
	require.Equal(t, int64(12+17+17), sum, "sum")

	sum = 0
	for _, i := range l {
		sum += i.Value
	}
	require.Equal(t, int64(12+17+17), sum, "sum2")

	// remove if

	odds := l.Copy()
	odds.RemoveIf(func(_ int, item *Item) bool { return item.Value%2 == 1 })
	assertDeepEqual(t, list.NewDeep(&Item{Value: 12}), odds, "wrong odds")

	// keep if

	evens := l.Copy()
	evens.KeepIf(func(_ int, item *Item) bool { return item.Value%2 == 1 })
	assertDeepEqual(t, list.NewDeep(&Item{Value: 17}, &Item{Value: 17}), evens, "wrong evens")

	// check source of copy

	assertDeepEqual(t, list.NewDeep(&Item{Value: 12}, &Item{Value: 17}, &Item{Value: 17}), l, "source of copy has been modified")

	// clear

	require.Equal(t, 3, len(l), "len before clear")
	require.False(t, l.IsEmpty(), "is_empty before clear")
	l.Clear()
	require.Equal(t, 0, len(l), "len after clear")
	require.True(t, l.IsEmpty(), "is_empty after clear")
}

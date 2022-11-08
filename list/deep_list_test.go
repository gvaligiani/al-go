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

	// list = 17, 12, 17 ( order re-shuffled during remove )

	// predicate

	require.True(t, l.AllOf(func(i *Item) bool { return i.Value < 20 }), "all_of")
	require.False(t, l.AllOf(func(i *Item) bool { return i.Value < 10 }), "all_of")
	require.True(t, l.NoneOf(func(i *Item) bool { return i.Value < 10 }), "none_of")
	require.False(t, l.AnyOf(func(i *Item) bool { return i.Value < 10 }), "any_of")

	// find

	index, found := l.FindIndex(&Item{Value: 12})
	require.Equal(t, 1, index, "index 12")
	require.True(t, found, "find 12")
	index, found = l.FindIndex(&Item{Value: 15})
	require.Equal(t, -1, index, "index 15")
	require.False(t, found, "find 15")

	item, found := l.FindIf(func(i *Item) bool { return i.Value%2 == 0 })
	require.Equal(t, &Item{Value: 12}, item, "odd item")
	require.True(t, found, "found odd")

	index, item, found = l.FindIndexIfNot(func(_ int, i *Item) bool { return i.Value%2 == 0 })
	require.Equal(t, &Item{Value: 17}, item, "even item")
	require.True(t, found, "found even")

	// range

	var sum int64

	sum = 0
	l.Each(func(i *Item) { sum += i.Value })
	require.Equal(t, int64(12+17+17), sum, "sum")

	sum = 0
	for _, i := range l {
		sum += i.Value
	}
	require.Equal(t, int64(12+17+17), sum, "sum2")

	// remove if

	odds := l.Copy()
	updated := odds.RemoveIf(func(item *Item) bool { return item.Value%2 == 1 })
	require.Equal(t, true, updated, "wrong updated!")
	assertDeepEqual(t, list.NewDeep(&Item{Value: 12}), odds, "wrong odds")

	// keep if

	evens := l.Copy()
	updated = evens.KeepIf(func(item *Item) bool { return item.Value%2 == 1 })
	require.Equal(t, true, updated, "wrong updated!")
	assertDeepEqual(t, list.NewDeep(&Item{Value: 17}, &Item{Value: 17}), evens, "wrong evens")

	// check source of copy

	assertDeepEqual(t, list.NewDeep(&Item{Value: 17}, &Item{Value: 12}, &Item{Value: 17}), l, "source of copy has been modified")

	// clear

	require.Equal(t, 3, len(l), "len before clear")
	require.False(t, l.IsEmpty(), "is_empty before clear")
	updated = l.Clear()
	require.Equal(t, true, updated, "wrong updated!")
	require.Equal(t, 0, len(l), "len after clear")
	require.True(t, l.IsEmpty(), "is_empty after clear")
	updated = l.Clear()
	require.Equal(t, false, updated, "wrong updated!")
}

package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gvaligiani/al.go/dict"
)

func TestDict(t *testing.T) {

	// builder

	d := dict.Dict[int, Item]{
		10: Item{Value: 21},
		20: Item{Value: 22},
	}

	// add

	require.False(t, d.Add(30, Item{Value: 15}), "add 30-15")
	require.True(t, d.Add(30, Item{Value: 17}), "override 30-17")

	// remove

	require.True(t, d.Remove(10), "remove 10")
	require.False(t, d.Remove(10), "remove 10 wice")

	// map[20:22,30:17]

	// predicate

	require.True(t, d.AllOf(func(i Item) bool { return i.Value < 30 }), "all_of")
	require.False(t, d.AllOf(func(i Item) bool { return i.Value < 10 }), "all_of")
	require.True(t, d.NoneOf(func(i Item) bool { return i.Value < 10 }), "none_of")
	require.False(t, d.AnyOf(func(i Item) bool { return i.Value < 10 }), "any_of")

	// find

	require.True(t, d.FindKey(30), "find 30")
	require.False(t, d.FindKey(50), "find 50")

	item, found := d.FindIf(func(i Item) bool { return i.Value%2 == 0 })
	require.Equal(t, Item{Value: 22}, item, "odd item")
	require.True(t, found, "found odd")

	item, found = d.FindIfNot(func(i Item) bool { return i.Value%2 == 0 })
	require.Equal(t, Item{Value: 17}, item, "even item")
	require.True(t, found, "found even")

	// range

	var sum int64

	sum = 0
	d.Each(func(i Item) { sum += i.Value })
	require.Equal(t, int64(22+17), sum, "sum - each")

	sum = 0
	for _, i := range d {
		sum += i.Value
	}
	require.Equal(t, int64(22+17), sum, "sum - range")

	// remove if

	odds := d.Copy()
	updated := odds.RemoveIf(func(item Item) bool { return item.Value%2 == 1 })
	require.Equal(t, true, updated, "wrong updated!")
	assertEqual(t, dict.Dict[int, Item]{20: {Value: 22}}, odds, "wrong odds")

	// keep if

	evens := d.Copy()
	updated = evens.KeepIf(func(item Item) bool { return item.Value%2 == 1 })
	require.Equal(t, true, updated, "wrong updated!")
	assertEqual(t, dict.Dict[int, Item]{30: {Value: 17}}, evens, "wrong evens")

	// check source of copy

	assertEqual(t, dict.Dict[int, Item]{20: {Value: 22}, 30: {Value: 17}}, d, "source of copy has been modified")

	// clear

	require.Equal(t, 2, len(d), "len before clear")
	require.False(t, d.IsEmpty(), "is_empty before clear")
	updated = d.Clear()
	require.Equal(t, true, updated, "wrong updated!")
	require.Equal(t, 0, len(d), "len after clear")
	require.True(t, d.IsEmpty(), "is_empty after clear")
	updated = d.Clear()
	require.Equal(t, false, updated, "wrong updated!")
}

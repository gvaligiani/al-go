package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gvaligiani/al.go/dict"
)

func TestAlias(t *testing.T) {

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

	require.True(t, d.AllOf(func(_ int, i Item) bool { return i.Value < 30 }), "all_of")
	require.False(t, d.AllOf(func(_ int, i Item) bool { return i.Value < 10 }), "all_of")
	require.True(t, d.NoneOf(func(_ int, i Item) bool { return i.Value < 10 }), "none_of")
	require.False(t, d.AnyOf(func(_ int, i Item) bool { return i.Value < 10 }), "any_of")

	// find

	require.True(t, d.FindKey(30), "find 30")
	require.False(t, d.FindKey(50), "find 50")

	key, item, found := d.FindIf(func(_ int, i Item) bool { return i.Value%2 == 0 })
	require.Equal(t, 20, key, "odd key")
	require.Equal(t, Item{Value: 22}, item, "odd item")
	require.True(t, found, "found odd")

	key, item, found = d.FindIfNot(func(_ int, i Item) bool { return i.Value%2 == 0 })
	require.Equal(t, 30, key, "even key")
	require.Equal(t, Item{Value: 17}, item, "even item")
	require.True(t, found, "found even")

	// range

	var sum int64

	sum = 0
	d.Each(func(_ int, i Item) { sum += i.Value })
	require.Equal(t, int64(22+17), sum, "sum - each")

	sum = 0
	for _, i := range d {
		sum += i.Value
	}
	require.Equal(t, int64(22+17), sum, "sum - range")

	// remove if

	odds := d.Copy()
	odds.RemoveIf(func(_ int, item Item) bool { return item.Value%2 == 1 })
	assertEquals(t, dict.Dict[int, Item]{20: {Value: 22}}, odds, "wrong odds")

	// keep if

	evens := d.Copy()
	evens.KeepIf(func(_ int, item Item) bool { return item.Value%2 == 1 })
	assertEquals(t, dict.Dict[int, Item]{30: {Value: 17}}, evens, "wrong evens")

	// check source of copy

	assertEquals(t, dict.Dict[int, Item]{20: {Value: 22}, 30: {Value: 17}}, d, "source of copy has been modified")

	// clear

	require.Equal(t, 2, len(d), "len before clear")
	require.False(t, d.IsEmpty(), "is_empty before clear")
	d.Clear()
	require.Equal(t, 0, len(d), "len after clear")
	require.True(t, d.IsEmpty(), "is_empty after clear")
}

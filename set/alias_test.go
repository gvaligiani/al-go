package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gvaligiani/algo/set"
)

func TestAlias(t *testing.T) {

	// builder

	s := set.New[Item](
		Item{Value:10},
		Item{Value:12},
		Item{Value:10}, // duplicate
	)

	// add

	require.True(t,s.Add(Item{Value:17}),"add 17")
	require.False(t,s.Add(Item{Value:17}),"add 17 twice")
	require.True(t,s.Add(Item{Value:11}),"add 11")

	// remove

	require.True(t,s.Remove(Item{Value:10}),"remove 10")
	require.False(t,s.Remove(Item{Value:10}),"remove 10 twice")
	require.True(t,s.Remove(Item{Value:17}),"remove 17")
	
	// predicate

	require.True(t,s.AllOf(func(i Item) bool{ return i.Value < 20}),"all_of")
	require.False(t,s.AllOf(func(i Item) bool{ return i.Value < 10}),"all_of")
	require.True(t,s.NoneOf(func(i Item) bool{ return i.Value < 10}),"none_of")
	require.False(t,s.AnyOf(func(i Item) bool{ return i.Value < 10}),"any_of")

	// find

	require.True(t,s.Find(Item{Value:11}),"find 11")
	require.False(t,s.Find(Item{Value:17}),"find 17")
	
	item, found := s.FindIf(func(i Item) bool { return i.Value%2 == 0 })
	require.Equal(t,Item{Value:12},item,"odd item")
	require.True(t,found,"found odd")
	
	item, found = s.FindIfNot(func(i Item) bool { return i.Value%2 == 0 })
	require.Equal(t,Item{Value:11},item,"even item")
	require.True(t,found,"found even")

	// range

	var sum int64

	sum = 0
	s.Each(func(i Item){ sum += i.Value })
	require.Equal(t,int64(11+12),sum,"sum")

	sum = 0
	for i := range s {
		sum += i.Value
	}
	require.Equal(t,int64(11+12),sum,"sum2")

	// clear

	require.Equal(t,2,len(s),"len before clear")
	s.Clear()
	require.Equal(t,0,len(s),"len after clear")
}

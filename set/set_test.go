package set_test

import (
	"testing"

	"github.com/gvaligiani/al-go/list"
	"github.com/gvaligiani/al-go/set"
)

func TestWrap(t *testing.T) {
	values := map[int]struct{}{1: {}, 3: {}, 2: {}}
	s := set.Wrap(values)
	set.Require(t, set.New(1, 2, 3), s)
}

func TestAdd(t *testing.T) {
	values := map[int]struct{}{1: {}, 3: {}, 2: {}}
	s := set.Wrap(values)
	s.Add(4)
	s.Add(3)
	set.Require(t, set.New(1, 2, 3, 4), s)
}

func TestDel(t *testing.T) {
	values := map[int]struct{}{1: {}, 3: {}, 2: {}}
	s := set.Wrap(values)
	s.Del(1)
	set.Require(t, set.New(2, 3), s)
}

func TestNew(t *testing.T) {
	s := set.New(1, 2, 1)
	set.Require(t, set.New(1, 2), s)
}

func TestToList(t *testing.T) {
	values := map[int]struct{}{1: {}, 3: {}, 2: {}}
	s := set.Wrap(values)
	list.RequireUnordered(t, list.New(1, 2, 3), s.ToList())
}

func TestToSorted(t *testing.T) {
	values := map[int]struct{}{1: {}, 3: {}, 2: {}}
	s := set.ToSorted(values)
	list.Require(t, list.New(1, 2, 3), s)
}

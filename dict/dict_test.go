package dict_test

import (
	"testing"

	"github.com/gvaligiani/al-go/dict"
	"github.com/gvaligiani/al-go/list"
	"github.com/gvaligiani/al-go/set"
)

func TestWrap(t *testing.T) {
	values := map[int]string{1: "one", 3: "three", 2: "two"}
	d := dict.Wrap(values)
	dict.Require(t, dict.New(dict.Pair(1, "one"), dict.Pair(2, "two"), dict.Pair(3, "three")), d)
}

func TestAdd(t *testing.T) {
	values := map[int]string{1: "one", 3: "three", 2: "two"}
	d := dict.Wrap(values)
	d.Add(4, "four")
	d.Add(3, "trois")
	dict.Require(t, dict.New(dict.Pair(1, "one"), dict.Pair(2, "two"), dict.Pair(3, "trois"), dict.Pair(4, "four")), d)
}

func TestDel(t *testing.T) {
	values := map[int]string{1: "one", 3: "three", 2: "two"}
	d := dict.Wrap(values)
	d.Del(1)
	dict.Require(t, dict.New(dict.Pair(2, "two"), dict.Pair(3, "three")), d)
}

func TestKeys(t *testing.T) {
	values := map[int]string{1: "one", 3: "three", 2: "two"}
	k := dict.Keys(values)
	list.RequireUnordered(t, list.New(3, 2, 1), k)
}

func TestSortedKeys(t *testing.T) {
	values := map[int]string{1: "one", 3: "three", 2: "two"}
	k := dict.SortedKeys(values)
	list.Require(t, list.New(1, 2, 3), k)
}

func TestKeySet(t *testing.T) {
	values := map[int]string{1: "one", 3: "three", 2: "two"}
	k := dict.KeySet(values)
	set.Require(t, set.New(3, 2, 1), k)
}

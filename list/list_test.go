package list_test

import (
	"testing"

	"github.com/gvaligiani/al-go/list"
)

func TestWrap(t *testing.T) {
	values := []int{1, 3, 2}
	l := list.Wrap(values)
	list.Require(t, list.New(1, 3, 2), l)
}

func TestAdd(t *testing.T) {
	values := []int{1, 3, 2}
	l := list.Wrap(values)
	l.Add(5, 4)
	list.Require(t, list.New(1, 3, 2, 5, 4), l)
}

func TestSort(t *testing.T) {
	values := []int{1, 3, 2, 3, 4}
	list.Sort(values)
	list.Require(t, list.New(1, 2, 3, 3, 4), values)
}

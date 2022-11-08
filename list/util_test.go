package list_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gvaligiani/al.go/list"
	"github.com/stretchr/testify/require"
)

// int64
var (
	EmptyInt64List   = list.New[int64]()
	DefaultInt64List = list.New[int64](
		21,
		12,
		34,
		87,
		52,
	)
	OtherInt64List = list.New[int64](
		21,
		12,
		87,
		52,
		69,
	)
)

// Item

type Item struct {
	Value int64
}

var (
	EmptyItemList   = list.New[Item]()
	DefaultItemList = list.New(
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 34},
		Item{Value: 87},
		Item{Value: 52},
	)
	OtherItemList = list.New(
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 87},
		Item{Value: 52},
		Item{Value: 69},
	)
)

// *Item

var (
	EmptyItemPointerList   = list.New[*Item]()
	DefaultItemPointerList = list.New(
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 34},
		&Item{Value: 87},
		&Item{Value: 52},
	)
	SameItemPointerList = list.New(
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 34},
		&Item{Value: 87},
		&Item{Value: 52},
	)
	OtherItemPointerList = list.New(
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 87},
		&Item{Value: 52},
		&Item{Value: 69},
	)
)

// assert

func assertEqual[V comparable, L ~[]V](t *testing.T, expected L, computed L, msg string) {
	require.Truef(t, list.Equal(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

func assertDeepEqual[V comparable, L ~[]V](t *testing.T, expected L, computed L, msg string) {
	require.Truef(t, list.DeepEqual(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

// dump

func toString[V comparable, L ~[]V](l L) string {
	if l == nil {
		return "<nil>"
	}
	var sb strings.Builder
	sb.WriteString("list[")
	first := true
	for _, v := range l {
		if first {
			sb.WriteString(fmt.Sprintf("%v", v))
		} else {
			sb.WriteString(fmt.Sprintf(" %v", v))
		}
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

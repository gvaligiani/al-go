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
	EmptyInt64List   = list.List[int64]{}
	DefaultInt64List = list.List[int64]{
		21,
		12,
		34,
		87,
		52,
	}
	OtherInt64List = list.List[int64]{
		21,
		12,
		87,
		52,
		69,
	}
)

// Item

type Item struct {
	Value int64
}

var (
	EmptyItemList   = list.List[Item]{}
	DefaultItemList = list.List[Item]{
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 34},
		Item{Value: 87},
		Item{Value: 52},
	}
	OtherItemList = list.List[Item]{
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 87},
		Item{Value: 52},
		Item{Value: 69},
	}
)

// *Item

var (
	EmptyItemPointerList   = list.List[*Item]{}
	DefaultItemPointerList = list.List[*Item]{
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 34},
		&Item{Value: 87},
		&Item{Value: 52},
	}
	OtherItemPointerList = list.List[*Item]{
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 87},
		&Item{Value: 52},
		&Item{Value: 69},
	}
)

// assert

func assertEquals[T comparable, L ~[]T](t *testing.T, expected L, computed L, msg string) {
	require.Truef(t, list.Equals(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

// dump

func toString[T comparable, L ~[]T](items L) string {
	if items == nil {
		return "<nil>"
	}
	var sb strings.Builder
	sb.WriteString("list[")
	first := true
	for _, item := range items {
		if first {
			sb.WriteString(fmt.Sprintf("%v", item))
		} else {
			sb.WriteString(fmt.Sprintf(" %v", item))
		}
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

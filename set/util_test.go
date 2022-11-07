package set_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gvaligiani/al.go/set"
)

// int64
var (
	EmptyInt64Set   = set.Set[int64]{}
	DefaultInt64Set = set.New[int64](
		21,
		12,
		34,
		87,
		52,
	)
	OtherInt64Set = set.New[int64](
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
	EmptyItemSet   = set.Set[Item]{}
	DefaultItemSet = set.New(
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 34},
		Item{Value: 87},
		Item{Value: 52},
	)
	OtherItemSet = set.New(
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 87},
		Item{Value: 52},
		Item{Value: 69},
	)
)

// *Item

var (
	EmptyItemPointerSet   = set.Set[*Item]{}
	DefaultItemPointerSet = set.New(
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 34},
		&Item{Value: 87},
		&Item{Value: 52},
	)
	SameItemPointerSet = set.New(
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 34},
		&Item{Value: 87},
		&Item{Value: 52},
	)
	OtherItemPointerSet = set.New(
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 87},
		&Item{Value: 52},
		&Item{Value: 69},
	)
)

// assert

func assertEqual[T comparable, S ~map[T]struct{}](t *testing.T, expected S, computed S, msg string) {
	require.Truef(t, set.Equal(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

func assertDeepEqual[T comparable, S ~map[T]struct{}](t *testing.T, expected S, computed S, msg string) {
	require.Truef(t, set.DeepEqual(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

// dump

func toString[T comparable, S ~map[T]struct{}](items S) string {
	if items == nil {
		return "<nil>"
	}
	var sb strings.Builder
	sb.WriteString("set[")
	first := true
	for item := range items {
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

package dict_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gvaligiani/al.go/dict"
)

// int64

var (
	EmptyInt64Dict   = dict.Dict[int, int64]{}
	DefaultInt64Dict = dict.Dict[int, int64]{
		10: 21,
		20: 12,
		30: 34,
		40: 87,
		50: 52,
	}
	OtherInt64Dict = dict.Dict[int, int64]{
		10: 21,
		20: 12,
		40: 87,
		50: 52,
		60: 69,
	}
)

// Item

type Item struct {
	Value int64
}

func (i *Item) String() string {
	if i == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ value: %d }", i.Value)
}

var (
	EmptyItemDict   = dict.Dict[int, Item]{}
	DefaultItemDict = dict.Dict[int, Item]{
		10: {Value: 21},
		20: {Value: 12},
		30: {Value: 34},
		40: {Value: 87},
		50: {Value: 52},
	}
	OtherItemDict = dict.Dict[int, Item]{
		10: {Value: 21},
		20: {Value: 12},
		40: {Value: 87},
		50: {Value: 52},
		60: {Value: 69},
	}
)

// *Item

var (
	EmptyItemPointerDict   = dict.Dict[int, *Item]{}
	DefaultItemPointerDict = dict.Dict[int, *Item]{
		10: {Value: 21},
		20: {Value: 12},
		30: {Value: 34},
		40: {Value: 87},
		50: {Value: 52},
	}
	SameItemPointerDict = dict.Dict[int, *Item]{
		10: {Value: 21},
		20: {Value: 12},
		30: {Value: 34},
		40: {Value: 87},
		50: {Value: 52},
	}
	OtherItemPointerDict = dict.Dict[int, *Item]{
		10: {Value: 21},
		20: {Value: 12},
		40: {Value: 87},
		50: {Value: 52},
		60: {Value: 69},
	}
)

// assert

func assertEqual[K comparable, T comparable](t *testing.T, expected map[K]T, computed map[K]T, msg string) {
	require.Truef(t, dict.Equal(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

func assertDeepEqual[K comparable, T any](t *testing.T, expected map[K]T, computed map[K]T, msg string) {
	require.Truef(t, dict.DeepEqual(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

// dump

func toString[K comparable, T any](items map[K]T) string {
	if items == nil {
		return "<nil>"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for key, item := range items {
		if first {
			sb.WriteString(fmt.Sprintf("%v:%v", key, item))
		} else {
			sb.WriteString(fmt.Sprintf(" %v:%v", key, item))
		}
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

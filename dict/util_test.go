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
	EmptyItemDict   = dict.New[int, Item]()
	DefaultItemDict = dict.New[int, Item]().
			With(10, Item{Value: 21}).
			With(20, Item{Value: 12}).
			With(30, Item{Value: 34}).
			With(40, Item{Value: 87}).
			With(50, Item{Value: 52})
	OtherItemDict = dict.New[int, Item]().
			With(10, Item{Value: 21}).
			With(20, Item{Value: 12}).
			With(40, Item{Value: 87}).
			With(50, Item{Value: 52}).
			With(60, Item{Value: 69})
)

// *Item

var (
	EmptyItemPointerDict   = dict.New[int, *Item]()
	DefaultItemPointerDict = dict.New[int, *Item]().
				With(10, &Item{Value: 21}).
				With(20, &Item{Value: 12}).
				With(30, &Item{Value: 34}).
				With(40, &Item{Value: 87}).
				With(50, &Item{Value: 52})

	SameItemPointerDict = dict.New[int, *Item]().
				With(10, &Item{Value: 21}).
				With(20, &Item{Value: 12}).
				With(30, &Item{Value: 34}).
				With(40, &Item{Value: 87}).
				With(50, &Item{Value: 52})

	OtherItemPointerDict = dict.New[int, *Item]().
				With(10, &Item{Value: 21}).
				With(20, &Item{Value: 12}).
				With(40, &Item{Value: 87}).
				With(50, &Item{Value: 52}).
				With(60, &Item{Value: 69})
)

// assert

func assertEqual[K comparable, V comparable, D ~map[K]V](t *testing.T, expected D, computed D, msg string) {
	require.Truef(t, dict.Equal(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

func assertDeepEqual[K comparable, V any, D ~map[K]V](t *testing.T, expected D, computed D, msg string) {
	require.Truef(t, dict.DeepEqual(expected, computed), "%s \n expected: %s \n computed: %s \n", msg, toString(expected), toString(computed))
}

// dump

func toString[K comparable, V any, D ~map[K]V](d D) string {
	if d == nil {
		return "<nil>"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range d {
		if first {
			sb.WriteString(fmt.Sprintf("%v:%v", k, v))
		} else {
			sb.WriteString(fmt.Sprintf(" %v:%v", k, v))
		}
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

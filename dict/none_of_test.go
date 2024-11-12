package dict_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al-go/dict"
	"github.com/gvaligiani/al-go/fn"
	"github.com/gvaligiani/al-go/test"
)

func TestNoneOf(t *testing.T) {

	type TestCase struct {
		dict      dict.Dict[string, int]
		predicate fn.BiPredicate[string, int]
		want      bool
	}

	testCases := map[string]TestCase{
		"nil": {
			dict:      nil,
			predicate: func(_ string, i int) bool { return i > 10 },
			want:      true,
		},
		"empty": {
			dict:      dict.New[string, int](),
			predicate: func(_ string, i int) bool { return i > 10 },
			want:      true,
		},
		"none": {
			dict:      dict.New(dict.Pair("A", 1), dict.Pair("B", 2), dict.Pair("C", 3)),
			predicate: func(_ string, i int) bool { return i > 10 },
			want:      true,
		},
		"one": {
			dict:      dict.New(dict.Pair("A", 1), dict.Pair("B", 2), dict.Pair("C", 3)),
			predicate: func(_ string, i int) bool { return i%2 == 0 },
			want:      false,
		},
		"some": {
			dict:      dict.New(dict.Pair("A", 1), dict.Pair("B", 2), dict.Pair("C", 3)),
			predicate: func(_ string, i int) bool { return i%2 == 1 },
			want:      false,
		},
		"all": {
			dict:      dict.New(dict.Pair("A", 1), dict.Pair("B", 2), dict.Pair("C", 3)),
			predicate: func(_ string, i int) bool { return i < 10 },
			want:      false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, ctx context.Context, logger *zap.Logger, tc TestCase) {

		got := dict.NoneOf(tc.dict, tc.predicate)
		require.Equal(t, tc.want, got, "wrong result")
	})
}

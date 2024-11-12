package list_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al-go/fn"
	"github.com/gvaligiani/al-go/list"
	"github.com/gvaligiani/al-go/test"
)

func TestNoneOf(t *testing.T) {

	type TestCase struct {
		list      list.List[int]
		predicate fn.Predicate[int]
		want      bool
	}

	testCases := map[string]TestCase{
		"nil": {
			list:      nil,
			predicate: func(i int) bool { return i > 10 },
			want:      true,
		},
		"empty": {
			list:      list.New[int](),
			predicate: func(i int) bool { return i > 10 },
			want:      true,
		},
		"none": {
			list:      list.New(1, 2, 3),
			predicate: func(i int) bool { return i > 10 },
			want:      true,
		},
		"one": {
			list:      list.New(1, 2, 3),
			predicate: func(i int) bool { return i%2 == 0 },
			want:      false,
		},
		"some": {
			list:      list.New(1, 2, 3),
			predicate: func(i int) bool { return i%2 == 1 },
			want:      false,
		},
		"all": {
			list:      list.New(1, 2, 3),
			predicate: func(i int) bool { return i < 10 },
			want:      false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, ctx context.Context, logger *zap.Logger, tc TestCase) {

		got := list.NoneOf(tc.list, tc.predicate)
		require.Equal(t, tc.want, got, "wrong result")
	})
}

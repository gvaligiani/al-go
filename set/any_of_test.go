package set_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al-go/fn"
	"github.com/gvaligiani/al-go/set"
	"github.com/gvaligiani/al-go/test"
)

func TestAnyOf(t *testing.T) {

	type TestCase struct {
		set       set.Set[int]
		predicate fn.Predicate[int]
		want      bool
	}

	testCases := map[string]TestCase{
		"nil": {
			set:       nil,
			predicate: func(i int) bool { return i > 10 },
			want:      false,
		},
		"empty": {
			set:       set.New[int](),
			predicate: func(i int) bool { return i > 10 },
			want:      false,
		},
		"none": {
			set:       set.New(1, 2, 3),
			predicate: func(i int) bool { return i > 10 },
			want:      false,
		},
		"one": {
			set:       set.New(1, 2, 3),
			predicate: func(i int) bool { return i%2 == 0 },
			want:      true,
		},
		"some": {
			set:       set.New(1, 2, 3),
			predicate: func(i int) bool { return i%2 == 1 },
			want:      true,
		},
		"all": {
			set:       set.New(1, 2, 3),
			predicate: func(i int) bool { return i < 10 },
			want:      true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, ctx context.Context, logger *zap.Logger, tc TestCase) {

		got := set.AnyOf(tc.set, tc.predicate)
		require.Equal(t, tc.want, got, "wrong result")
	})
}

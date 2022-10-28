package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestFindInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int64]struct{}
		value     int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     34,
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyInt64Set,
			value:     34,
			wantFound: false,
		},
		"found": {
			items:     test.Int64Set,
			value:     34,
			wantFound: true,
		},
		"not-found": {
			items:     test.Int64Set,
			value:     34,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find[int64](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[test.Item]struct{}
		value     test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     test.Item{Value: 34},
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemSet,
			value:     test.Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     test.ItemSet,
			value:     test.Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     test.ItemSet,
			value:     test.Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find[test.Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[*test.Item]struct{}
		value     *test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     &test.Item{Value: 34},
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemPointerSet,
			value:     &test.Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     test.ItemPointerSet,
			value:     &test.Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     test.ItemPointerSet,
			value:     &test.Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find[*test.Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

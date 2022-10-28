package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestAnyOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		predicate func(int, int64) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.Int64Dict,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.Int64Dict,
			predicate: func(_ int, i int64) bool { return i > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.Int64Dict,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf[int,int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]test.Item
		predicate func(int, test.Item) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.ItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.ItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.ItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf[int, test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]*test.Item
		predicate func(int, *test.Item) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.ItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.ItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.ItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf[int, *test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

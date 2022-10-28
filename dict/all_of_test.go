package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestAllOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		predicate func(int, int64) bool
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     test.EmptyInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     test.DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     test.DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     test.DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := dict.AllOf[int,int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]test.Item
		predicate func(int,test.Item) bool
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item test.Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     test.EmptyItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     test.DefaultItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     test.DefaultItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     test.DefaultItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := dict.AllOf[int,test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]*test.Item
		predicate func(int,*test.Item) bool
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     test.EmptyItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     test.DefaultItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     test.DefaultItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     test.DefaultItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := dict.AllOf[int, *test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

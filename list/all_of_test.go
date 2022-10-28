package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestAllOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		predicate func(int64) bool
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     test.EmptyInt64List,
			predicate: func(i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     test.DefaultInt64List,
			predicate: func(i int64) bool { return i > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     test.DefaultInt64List,
			predicate: func(i int64) bool { return i > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     test.DefaultInt64List,
			predicate: func(i int64) bool { return i < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := list.AllOf[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []test.Item
		predicate func(test.Item) bool
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     test.EmptyItemList,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     test.DefaultItemList,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     test.DefaultItemList,
			predicate: func(item test.Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     test.DefaultItemList,
			predicate: func(item test.Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := list.AllOf[test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []*test.Item
		predicate func(*test.Item) bool
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     test.EmptyItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     test.DefaultItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     test.DefaultItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     test.DefaultItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := list.AllOf[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

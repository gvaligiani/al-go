package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestAnyOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		predicate func(int64) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyInt64List,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.Int64List,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.Int64List,
			predicate: func(i int64) bool { return i > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.Int64List,
			predicate: func(i int64) bool { return i < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []test.Item
		predicate func(test.Item) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyItemList,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.ItemList,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.ItemList,
			predicate: func(item test.Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.ItemList,
			predicate: func(item test.Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf[test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []*test.Item
		predicate func(*test.Item) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.ItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.ItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.ItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

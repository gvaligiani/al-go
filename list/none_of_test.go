package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestNoneOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      []int64
		predicate  func(int64) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyInt64List,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.Int64List,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.Int64List,
			predicate:  func(i int64) bool { return i > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.Int64List,
			predicate:  func(i int64) bool { return i < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := list.NoneOf[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      []test.Item
		predicate  func(test.Item) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyItemList,
			predicate:  func(item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.ItemList,
			predicate:  func(item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.ItemList,
			predicate:  func(item test.Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.ItemList,
			predicate:  func(item test.Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := list.NoneOf[test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      []*test.Item
		predicate  func(*test.Item) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyItemPointerList,
			predicate:  func(item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.ItemPointerList,
			predicate:  func(item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.ItemPointerList,
			predicate:  func(item *test.Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.ItemPointerList,
			predicate:  func(item *test.Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := list.NoneOf[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

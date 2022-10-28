package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestNoneOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      map[int]int64
		predicate  func(int, int64) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyInt64Dict,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.Int64Dict,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.Int64Dict,
			predicate:  func(_ int, i int64) bool { return i > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.Int64Dict,
			predicate:  func(_ int, i int64) bool { return i < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := dict.NoneOf[int,int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      map[int]test.Item
		predicate  func(int, test.Item) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(_ int, item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyItemDict,
			predicate:  func(_ int, item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.ItemDict,
			predicate:  func(_ int, item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.ItemDict,
			predicate:  func(_ int, item test.Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.ItemDict,
			predicate:  func(_ int, item test.Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := dict.NoneOf[int,test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      map[int]*test.Item
		predicate  func(int, *test.Item) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyItemPointerDict,
			predicate:  func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.ItemPointerDict,
			predicate:  func(_ int, item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.ItemPointerDict,
			predicate:  func(_ int, item *test.Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.ItemPointerDict,
			predicate:  func(_ int, item *test.Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := dict.NoneOf[int,*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

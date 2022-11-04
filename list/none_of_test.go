package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestNoneOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      []int64
		predicate  list.Predicate[int64]
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      EmptyInt64List,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      DefaultInt64List,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      DefaultInt64List,
			predicate:  func(_ int, i int64) bool { return i > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      DefaultInt64List,
			predicate:  func(_ int, i int64) bool { return i < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := list.NoneOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      list.List[Item]
		predicate  list.Predicate[Item]
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(_ int, item Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      EmptyItemList,
			predicate:  func(_ int, item Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      DefaultItemList,
			predicate:  func(_ int, item Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      DefaultItemList,
			predicate:  func(_ int, item Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      DefaultItemList,
			predicate:  func(_ int, item Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := list.NoneOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      list.List[*Item]
		predicate  list.Predicate[*Item]
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(_ int, item *Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      EmptyItemPointerList,
			predicate:  func(_ int, item *Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      DefaultItemPointerList,
			predicate:  func(_ int, item *Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      DefaultItemPointerList,
			predicate:  func(_ int, item *Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      DefaultItemPointerList,
			predicate:  func(_ int, item *Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := list.NoneOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

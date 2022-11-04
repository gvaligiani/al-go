package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestAllOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		predicate list.Predicate[int64]
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyInt64List,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := list.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[Item]
		predicate list.Predicate[Item]
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyItemList,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := list.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[*Item]
		predicate list.Predicate[*Item]
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := list.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

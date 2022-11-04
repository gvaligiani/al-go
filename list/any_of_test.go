package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestAnyOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[int64]
		predicate list.Predicate[int64]
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyInt64List,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[Item]
		predicate list.Predicate[Item]
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyItemList,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[*Item]
		predicate list.Predicate[*Item]
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

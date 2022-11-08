package dict_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestCopyIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, int64]
		predicate dict.Predicate[int, int64]
		wantItems dict.Dict[int, int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyInt64Dict,
			predicate: func(_ int, i int64) bool { return i%2 == 0 },
			wantItems: EmptyInt64Dict,
		},
		"copy-none": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return false },
			wantItems: EmptyInt64Dict,
		},
		"copy-odd": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i%2 == 0 },
			wantItems: dict.Dict[int, int64]{
				20: 12,
				30: 34,
				50: 52,
			},
		},
		"copy-even": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i%2 == 1 },
			wantItems: dict.Dict[int, int64]{
				10: 21,
				40: 87,
			},
		},
		"copy-all": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return true },
			wantItems: DefaultInt64Dict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, Item]
		predicate dict.Predicate[int, Item]
		wantItems dict.Dict[int, Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemDict,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemDict,
		},
		"copy-none": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return false },
			wantItems: EmptyItemDict,
		},
		"copy-odd": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantItems: dict.Dict[int, Item]{
				20: {Value: 12},
				30: {Value: 34},
				50: {Value: 52},
			},
		},
		"copy-even": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 1 },
			wantItems: dict.Dict[int, Item]{
				10: {Value: 21},
				40: {Value: 87},
			},
		},
		"copy-all": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return true },
			wantItems: DefaultItemDict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, *Item]
		predicate dict.Predicate[int, *Item]
		wantItems dict.Dict[int, *Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemPointerDict,
		},
		"copy-none": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return false },
			wantItems: EmptyItemPointerDict,
		},
		"copy-odd": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantItems: dict.Dict[int, *Item]{
				20: {Value: 12},
				30: {Value: 34},
				50: {Value: 52},
			},
		},
		"copy-even": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 1 },
			wantItems: dict.Dict[int, *Item]{
				10: {Value: 21},
				40: {Value: 87},
			},
		},
		"copy-all": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return true },
			wantItems: DefaultItemPointerDict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

package dict_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestKeepIfInt64(t *testing.T) {

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
		"keep-none": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return false },
			wantItems: EmptyInt64Dict,
		},
		"keep-odd": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i%2 == 0 },
			wantItems: dict.Dict[int, int64]{
				20: 12,
				30: 34,
				50: 52,
			},
		},
		"keep-even": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i%2 == 1 },
			wantItems: dict.Dict[int, int64]{
				10: 21,
				40: 87,
			},
		},
		"keep-all": {
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
		gotItems := dict.Copy(testCase.items)
		dict.KeepIf(&gotItems, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestKeepIfStruct(t *testing.T) {

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
		"keep-none": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return false },
			wantItems: EmptyItemDict,
		},
		"keep-odd": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantItems: dict.Dict[int, Item]{
				20: {Value: 12},
				30: {Value: 34},
				50: {Value: 52},
			},
		},
		"keep-even": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 1 },
			wantItems: dict.Dict[int, Item]{
				10: {Value: 21},
				40: {Value: 87},
			},
		},
		"keep-all": {
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
		gotItems := dict.Copy(testCase.items)
		dict.KeepIf(&gotItems, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestKeepIfStructPointer(t *testing.T) {

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
		"keep-none": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return false },
			wantItems: EmptyItemPointerDict,
		},
		"keep-odd": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantItems: dict.Dict[int, *Item]{
				20: {Value: 12},
				30: {Value: 34},
				50: {Value: 52},
			},
		},
		"keep-even": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 1 },
			wantItems: dict.Dict[int, *Item]{
				10: {Value: 21},
				40: {Value: 87},
			},
		},
		"keep-all": {
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
		gotItems := dict.Copy(testCase.items)
		dict.KeepIf(&gotItems, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

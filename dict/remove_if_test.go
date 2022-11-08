package dict_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
	"github.com/stretchr/testify/require"
)

func TestRemoveIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       dict.Dict[int, int64]
		predicate   dict.Predicate[int, int64]
		wantUpdated bool
		wantItems   dict.Dict[int, int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(_ int, i int64) bool { return i%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyInt64Dict,
			predicate:   func(_ int, i int64) bool { return i%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyInt64Dict,
		},
		"remove-none": {
			items:       DefaultInt64Dict,
			predicate:   func(_ int, i int64) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultInt64Dict,
		},
		"remove-odd": {
			items:       DefaultInt64Dict,
			predicate:   func(_ int, i int64) bool { return i%2 == 0 },
			wantUpdated: true,
			wantItems: dict.Dict[int, int64]{
				10: 21,
				40: 87,
			},
		},
		"remove-even": {
			items:       DefaultInt64Dict,
			predicate:   func(_ int, i int64) bool { return i%2 == 1 },
			wantUpdated: true,
			wantItems: dict.Dict[int, int64]{
				20: 12,
				30: 34,
				50: 52,
			},
		},
		"remove-all": {
			items:       DefaultInt64Dict,
			predicate:   func(_ int, i int64) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyInt64Dict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.Copy(testCase.items)
		gotUpdated := dict.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestRemoveIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       dict.Dict[int, Item]
		predicate   dict.Predicate[int, Item]
		wantUpdated bool
		wantItems   dict.Dict[int, Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyItemDict,
			predicate:   func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyItemDict,
		},
		"remove-none": {
			items:       DefaultItemDict,
			predicate:   func(_ int, item Item) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultItemDict,
		},
		"remove-odd": {
			items:       DefaultItemDict,
			predicate:   func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: dict.Dict[int, Item]{
				10: {Value: 21},
				40: {Value: 87},
			},
		},
		"remove-even": {
			items:       DefaultItemDict,
			predicate:   func(_ int, item Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: dict.Dict[int, Item]{
				20: {Value: 12},
				30: {Value: 34},
				50: {Value: 52},
			},
		},
		"remove-all": {
			items:       DefaultItemDict,
			predicate:   func(_ int, item Item) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyItemDict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.Copy(testCase.items)
		gotUpdated := dict.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestRemoveIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       dict.Dict[int, *Item]
		predicate   dict.Predicate[int, *Item]
		wantUpdated bool
		wantItems   dict.Dict[int, *Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyItemPointerDict,
			predicate:   func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyItemPointerDict,
		},
		"remove-none": {
			items:       DefaultItemPointerDict,
			predicate:   func(_ int, item *Item) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultItemPointerDict,
		},
		"remove-odd": {
			items:       DefaultItemPointerDict,
			predicate:   func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: dict.Dict[int, *Item]{
				10: {Value: 21},
				40: {Value: 87},
			},
		},
		"remove-even": {
			items:       DefaultItemPointerDict,
			predicate:   func(_ int, item *Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: dict.Dict[int, *Item]{
				20: {Value: 12},
				30: {Value: 34},
				50: {Value: 52},
			},
		},
		"remove-all": {
			items:       DefaultItemPointerDict,
			predicate:   func(_ int, item *Item) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyItemPointerDict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.Copy(testCase.items)
		gotUpdated := dict.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

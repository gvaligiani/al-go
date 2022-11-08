package list_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
	"github.com/gvaligiani/al.go/util"
	"github.com/stretchr/testify/require"
)

func TestKeepIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       list.List[int64]
		predicate   util.Predicate[int64]
		wantUpdated bool
		wantItems   list.List[int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(i int64) bool { return i%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyInt64List,
			predicate:   func(i int64) bool { return i%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyInt64List,
		},
		"keep-none": {
			items:       DefaultInt64List,
			predicate:   func(i int64) bool { return false },
			wantUpdated: true,
			wantItems:   EmptyInt64List,
		},
		"keep-odd": {
			items:       DefaultInt64List,
			predicate:   func(i int64) bool { return i%2 == 0 },
			wantUpdated: true,
			wantItems: list.New[int64]( // list re-shuffled
				52,
				12,
				34,
			),
		},
		"keep-even": {
			items:       DefaultInt64List,
			predicate:   func(i int64) bool { return i%2 == 1 },
			wantUpdated: true,
			wantItems: list.New[int64](
				21,
				87,
			),
		},
		"keep-all": {
			items:       DefaultInt64List,
			predicate:   func(i int64) bool { return true },
			wantUpdated: false,
			wantItems:   DefaultInt64List,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		gotUpdated := list.KeepIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestKeepIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       list.List[Item]
		predicate   util.Predicate[Item]
		wantUpdated bool
		wantItems   list.List[Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(item Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyItemList,
			predicate:   func(item Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyItemList,
		},
		"keep-none": {
			items:       DefaultItemList,
			predicate:   func(item Item) bool { return false },
			wantUpdated: true,
			wantItems:   EmptyItemList,
		},
		"keep-odd": {
			items:       DefaultItemList,
			predicate:   func(item Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: list.New( // list re-shuffled
				Item{Value: 52},
				Item{Value: 12},
				Item{Value: 34},
			),
		},
		"keep-even": {
			items:       DefaultItemList,
			predicate:   func(item Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: list.New(
				Item{Value: 21},
				Item{Value: 87},
			),
		},
		"keep-all": {
			items:       DefaultItemList,
			predicate:   func(item Item) bool { return true },
			wantUpdated: false,
			wantItems:   DefaultItemList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		gotUpdated := list.KeepIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestKeepIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       list.List[*Item]
		predicate   util.Predicate[*Item]
		wantUpdated bool
		wantItems   list.List[*Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyItemPointerList,
			predicate:   func(item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyItemPointerList,
		},
		"keep-none": {
			items:       DefaultItemPointerList,
			predicate:   func(item *Item) bool { return false },
			wantUpdated: true,
			wantItems:   EmptyItemPointerList,
		},
		"keep-odd": {
			items:       DefaultItemPointerList,
			predicate:   func(item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: list.New( // list re-shuffled
				&Item{Value: 52},
				&Item{Value: 12},
				&Item{Value: 34},
			),
		},
		"keep-even": {
			items:       DefaultItemPointerList,
			predicate:   func(item *Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: list.New(
				&Item{Value: 21},
				&Item{Value: 87},
			),
		},
		"keep-all": {
			items:       DefaultItemPointerList,
			predicate:   func(item *Item) bool { return true },
			wantUpdated: false,
			wantItems:   DefaultItemPointerList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		gotUpdated := list.KeepIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

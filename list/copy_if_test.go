package list_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestCopyIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[int64]
		predicate list.Predicate[int64]
		wantItems list.List[int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyInt64List,
			predicate: func(_ int, i int64) bool { return i%2 == 0 },
			wantItems: EmptyInt64List,
		},
		"copy-none": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return false },
			wantItems: EmptyInt64List,
		},
		"copy-odd": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%2 == 0 },
			wantItems: list.New[int64](
				12,
				34,
				52,
			),
		},
		"copy-even": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%2 == 1 },
			wantItems: list.New[int64](
				21,
				87,
			),
		},
		"copy-all": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return true },
			wantItems: DefaultInt64List,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[Item]
		predicate list.Predicate[Item]
		wantItems list.List[Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemList,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemList,
		},
		"copy-none": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return false },
			wantItems: EmptyItemList,
		},
		"copy-odd": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantItems: list.New(
				Item{Value: 12},
				Item{Value: 34},
				Item{Value: 52},
			),
		},
		"copy-even": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 1 },
			wantItems: list.New(
				Item{Value: 21},
				Item{Value: 87},
			),
		},
		"copy-all": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return true },
			wantItems: DefaultItemList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[*Item]
		predicate list.Predicate[*Item]
		wantItems list.List[*Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemPointerList,
		},
		"copy-none": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return false },
			wantItems: EmptyItemPointerList,
		},
		"copy-odd": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantItems: list.New(
				&Item{Value: 12},
				&Item{Value: 34},
				&Item{Value: 52},
			),
		},
		"copy-even": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 1 },
			wantItems: list.New(
				&Item{Value: 21},
				&Item{Value: 87},
			),
		},
		"copy-all": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return true },
			wantItems: DefaultItemPointerList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

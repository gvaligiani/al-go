package list_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestCopyIfNotInt64(t *testing.T) {

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
		"remove-none": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return false },
			wantItems: DefaultInt64List,
		},
		"remove-odd": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%2 == 0 },
			wantItems: list.New[int64](
				21,
				87,
			),
		},
		"remove-even": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%2 == 1 },
			wantItems: list.New[int64](
				12,
				34,
				52,
			),
		},
		"remove-all": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return true },
			wantItems: EmptyInt64List,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.CopyIfNot(testCase.items, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfNotStruct(t *testing.T) {

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
		"remove-none": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return false },
			wantItems: DefaultItemList,
		},
		"remove-odd": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantItems: list.New(
				Item{Value: 21},
				Item{Value: 87},
			),
		},
		"remove-even": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%2 == 1 },
			wantItems: list.New(
				Item{Value: 12},
				Item{Value: 34},
				Item{Value: 52},
			),
		},
		"remove-all": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return true },
			wantItems: EmptyItemList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.CopyIfNot(testCase.items, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfNotStructPointer(t *testing.T) {

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
		"remove-none": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return false },
			wantItems: DefaultItemPointerList,
		},
		"remove-odd": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantItems: list.New(
				&Item{Value: 21},
				&Item{Value: 87},
			),
		},
		"remove-even": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%2 == 1 },
			wantItems: list.New(
				&Item{Value: 12},
				&Item{Value: 34},
				&Item{Value: 52},
			),
		},
		"remove-all": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return true },
			wantItems: EmptyItemPointerList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.CopyIfNot(testCase.items, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

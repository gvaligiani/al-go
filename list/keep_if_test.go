package list_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestKeepIfInt64(t *testing.T) {

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
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyInt64List,
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: EmptyInt64List,
		},
		"keep-none": {
			items:     DefaultInt64List,
			predicate: func(i int64) bool { return false },
			wantItems: EmptyInt64List,
		},
		"keep-odd": {
			items:     DefaultInt64List,
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: list.New[int64]( // list re-shuffled
				52,
				12,
				34,
			),
		},
		"keep-even": {
			items:     DefaultInt64List,
			predicate: func(i int64) bool { return i%2 == 1 },
			wantItems: list.New[int64](
				21,
				87,
			),
		},
		"keep-all": {
			items:     DefaultInt64List,
			predicate: func(i int64) bool { return true },
			wantItems: DefaultInt64List,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		list.KeepIf(&gotItems, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestKeepIfStruct(t *testing.T) {

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
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemList,
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemList,
		},
		"keep-none": {
			items:     DefaultItemList,
			predicate: func(item Item) bool { return false },
			wantItems: EmptyItemList,
		},
		"keep-odd": {
			items:     DefaultItemList,
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: list.New( // list re-shuffled
				Item{Value: 52},
				Item{Value: 12},
				Item{Value: 34},
			),
		},
		"keep-even": {
			items:     DefaultItemList,
			predicate: func(item Item) bool { return item.Value%2 == 1 },
			wantItems: list.New(
				Item{Value: 21},
				Item{Value: 87},
			),
		},
		"keep-all": {
			items:     DefaultItemList,
			predicate: func(item Item) bool { return true },
			wantItems: DefaultItemList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		list.KeepIf(&gotItems, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestKeepIfStructPointer(t *testing.T) {

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
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemPointerList,
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemPointerList,
		},
		"keep-none": {
			items:     DefaultItemPointerList,
			predicate: func(item *Item) bool { return false },
			wantItems: EmptyItemPointerList,
		},
		"keep-odd": {
			items:     DefaultItemPointerList,
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: list.New( // list re-shuffled
				&Item{Value: 52},
				&Item{Value: 12},
				&Item{Value: 34},
			),
		},
		"keep-even": {
			items:     DefaultItemPointerList,
			predicate: func(item *Item) bool { return item.Value%2 == 1 },
			wantItems: list.New(
				&Item{Value: 21},
				&Item{Value: 87},
			),
		},
		"keep-all": {
			items:     DefaultItemPointerList,
			predicate: func(item *Item) bool { return true },
			wantItems: DefaultItemPointerList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		list.KeepIf(&gotItems, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

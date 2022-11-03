package set_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
)

func TestCopyIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[int64]
		predicate set.Predicate[int64]
		wantItems set.Set[int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyInt64Set,
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: EmptyInt64Set,
		},
		"copy-none": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return false },
			wantItems: EmptyInt64Set,
		},
		"copy-odd": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: set.New[int64](
				12,
				34,
				52,
			),
		},
		"copy-even": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i%2 == 1 },
			wantItems: set.New[int64](
				21,
				87,
			),
		},
		"copy-all": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return true },
			wantItems: DefaultInt64Set,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[Item]
		predicate func(Item) bool
		wantItems set.Set[Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemSet,
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemSet,
		},
		"copy-none": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return false },
			wantItems: EmptyItemSet,
		},
		"copy-odd": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: set.New(
				Item{Value: 12},
				Item{Value: 34},
				Item{Value: 52},
			),
		},
		"copy-even": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value%2 == 1 },
			wantItems: set.New(
				Item{Value: 21},
				Item{Value: 87},
			),
		},
		"copy-all": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return true },
			wantItems: DefaultItemSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[*Item]
		predicate func(*Item) bool
		wantItems set.Set[*Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemPointerSet,
		},
		"copy-none": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return false },
			wantItems: EmptyItemPointerSet,
		},
		"copy-odd": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: set.New(
				&Item{Value: 12},
				&Item{Value: 34},
				&Item{Value: 52},
			),
		},
		"copy-even": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%2 == 1 },
			wantItems: set.New(
				&Item{Value: 21},
				&Item{Value: 87},
			),
		},
		"copy-all": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return true },
			wantItems: DefaultItemPointerSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.CopyIf(testCase.items, testCase.predicate)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

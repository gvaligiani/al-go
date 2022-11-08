package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
	"github.com/gvaligiani/al.go/util"
)

func TestFindIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		predicate util.Predicate[int64]
		wantIndex int
		wantItem  int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"empty": {
			items:     EmptyInt64List,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     DefaultInt64List,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"one-match": {
			items:     DefaultInt64List,
			predicate: func(i int64) bool { return i%10 == 4 },
			wantItem:  34,
			wantFound: true,
		},
		"two-matches": {
			items:     DefaultInt64List,
			predicate: func(i int64) bool { return i%10 == 2 },
			wantItem:  12,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[Item]
		predicate util.Predicate[Item]
		wantItem  Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item Item) bool { return item.Value%10 == 3 },
			wantItem:  Item{},
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemList,
			predicate: func(item Item) bool { return item.Value%10 == 3 },
			wantItem:  Item{},
			wantFound: false,
		},
		"no-match": {
			items:     DefaultItemList,
			predicate: func(item Item) bool { return item.Value%10 == 3 },
			wantItem:  Item{},
			wantFound: false,
		},
		"one-match": {
			items:     DefaultItemList,
			predicate: func(item Item) bool { return item.Value%10 == 4 },
			wantItem:  Item{Value: 34},
			wantFound: true,
		},
		"two-matches": {
			items:     DefaultItemList,
			predicate: func(item Item) bool { return item.Value%10 == 2 },
			wantItem:  Item{Value: 12},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[*Item]
		predicate util.Predicate[*Item]
		wantItem  *Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemPointerList,
			predicate: func(item *Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     DefaultItemPointerList,
			predicate: func(item *Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"one-match": {
			items:     DefaultItemPointerList,
			predicate: func(item *Item) bool { return item.Value%10 == 4 },
			wantItem:  &Item{Value: 34},
			wantFound: true,
		},
		"two-matches": {
			items:     DefaultItemPointerList,
			predicate: func(item *Item) bool { return item.Value%10 == 2 },
			wantItem:  &Item{Value: 12},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIndexIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		predicate util.BiPredicate[int, int64]
		wantIndex int
		wantItem  int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantIndex: -1,
			wantItem:  0,
			wantFound: false,
		},
		"empty": {
			items:     EmptyInt64List,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantIndex: -1,
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantIndex: -1,
			wantItem:  0,
			wantFound: false,
		},
		"one-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%10 == 4 },
			wantIndex: 2,
			wantItem:  34,
			wantFound: true,
		},
		"two-matches": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%10 == 2 },
			wantIndex: 1,
			wantItem:  12,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotItem, gotFound := list.FindIndexIf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIndexIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[Item]
		predicate util.BiPredicate[int, Item]
		wantIndex int
		wantItem  Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantItem:  Item{},
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemList,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantItem:  Item{},
			wantFound: false,
		},
		"no-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantItem:  Item{},
			wantFound: false,
		},
		"one-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 4 },
			wantIndex: 2,
			wantItem:  Item{Value: 34},
			wantFound: true,
		},
		"two-matches": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 2 },
			wantIndex: 1,
			wantItem:  Item{Value: 12},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotItem, gotFound := list.FindIndexIf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIndexIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[*Item]
		predicate util.BiPredicate[int, *Item]
		wantIndex int
		wantItem  *Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantItem:  nil,
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantItem:  nil,
			wantFound: false,
		},
		"one-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 4 },
			wantIndex: 2,
			wantItem:  &Item{Value: 34},
			wantFound: true,
		},
		"two-matches": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 2 },
			wantIndex: 1,
			wantItem:  &Item{Value: 12},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotItem, gotFound := list.FindIndexIf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

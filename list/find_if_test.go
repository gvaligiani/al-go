package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestFindIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		predicate func(int64) bool
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
			items:     test.EmptyInt64List,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     test.Int64List,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"one-match": {
			items:     test.Int64List,
			predicate: func(i int64) bool { return i%10 == 4 },
			wantItem:  34,
			wantFound: true,
		},
		"two-matches": {
			items:     test.Int64List,
			predicate: func(i int64) bool { return i%10 == 2 },
			wantItem:  12,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIf[int64](testCase.items, testCase.predicate)

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
		items     []test.Item
		predicate func(test.Item) bool
		wantItem  test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item test.Item) bool { return item.Value%10 == 3 },
			wantItem:  test.Item{},
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemList,
			predicate: func(item test.Item) bool { return item.Value%10 == 3 },
			wantItem:  test.Item{},
			wantFound: false,
		},
		"no-match": {
			items:     test.ItemList,
			predicate: func(item test.Item) bool { return item.Value%10 == 3 },
			wantItem:  test.Item{},
			wantFound: false,
		},
		"one-match": {
			items:     test.ItemList,
			predicate: func(item test.Item) bool { return item.Value%10 == 4 },
			wantItem:  test.Item{Value: 34},
			wantFound: true,
		},
		"two-matches": {
			items:     test.ItemList,
			predicate: func(item test.Item) bool { return item.Value%10 == 2 },
			wantItem:  test.Item{Value: 12},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIf[test.Item](testCase.items, testCase.predicate)

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
		items     []*test.Item
		predicate func(*test.Item) bool
		wantItem  *test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *test.Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     test.ItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"one-match": {
			items:     test.ItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value%10 == 4 },
			wantItem:  &test.Item{Value: 34},
			wantFound: true,
		},
		"two-matches": {
			items:     test.ItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value%10 == 2 },
			wantItem:  &test.Item{Value: 12},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIf[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

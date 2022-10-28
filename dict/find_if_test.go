package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestFindIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		predicate func(int, int64) bool
		wantKey   int
		wantItem  int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantKey: 0,
			wantItem:  0,
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyInt64Dict,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantKey: 0,
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     test.DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantKey: 0,
			wantItem:  0,
			wantFound: false,
		},
		"one-match": {
			items:     test.DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i%10 == 4 },
			wantKey: 30,
			wantItem:  34,
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     test.DefaultInt64Dict,
		// 	predicate: func(_ int, i int64) bool { return i%10 == 2 },
		// 	wantKey:   20,
		// 	wantItem:  12,
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIf[int,int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]test.Item
		predicate func(int, test.Item) bool
		wantKey  int
		wantItem  test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item test.Item) bool { return item.Value%10 == 3 },
			wantKey: 0,
			wantItem:  test.Item{},
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value%10 == 3 },
			wantKey: 0,
			wantItem:  test.Item{},
			wantFound: false,
		},
		"no-match": {
			items:     test.DefaultItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value%10 == 3 },
			wantKey: 0,
			wantItem:  test.Item{},
			wantFound: false,
		},
		"one-match": {
			items:     test.DefaultItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value%10 == 4 },
			wantKey: 30,
			wantItem:  test.Item{Value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     test.ItemDict,
		// 	predicate: func(_ int, item test.Item) bool { return item.Value%10 == 2 },
		// 	wantKey:   20,
		// 	wantItem:  test.Item{Value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIf[int, test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]*test.Item
		predicate func(int, *test.Item) bool
		wantKey int
		wantItem  *test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *test.Item) bool { return item.Value%10 == 3 },
			wantKey: 0,
			wantItem:  nil,
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value%10 == 3 },
			wantKey: 0,
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     test.DefaultItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value%10 == 3 },
			wantKey: 0,
			wantItem:  nil,
			wantFound: false,
		},
		"one-match": {
			items:     test.DefaultItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value%10 == 4 },
			wantKey: 30,
			wantItem:  &test.Item{Value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     test.ItemPointerDict,
		// 	predicate: func(_ int, item *test.Item) bool { return item.Value%10 == 2 },
		// 	wantKey:   20,
		// 	wantItem:  &test.Item{Value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIf[int, *test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

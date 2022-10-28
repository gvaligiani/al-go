package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestFindIfNotInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		predicate func(int, int64) bool
		wantKey int
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
			items:     test.Int64Dict,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantKey: 0,
			wantItem:  0,
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     test.Int64Dict,
		// 	predicate: func(_ int, i int64) bool { return i%10 == 4 },
		// 	wantKey:   10,
		// 	wantItem:  21,
		// 	wantFound: true,
		// },
		"one-match": {
			items:     test.Int64Dict,
			predicate: func(_ int, i int64) bool { return i < 60 },
			wantKey: 40,
			wantItem:  87,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIfNot[int, int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindIfNotStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]test.Item
		predicate func(int, test.Item) bool
		wantKey int
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
			items:     test.ItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value < 100 },
			wantKey: 0,
			wantItem:  test.Item{},
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     test.ItemDict,
		// 	predicate: func(_ int, item test.Item) bool { return item.Value%10 == 4 },
		// 	wantKey:   10,
		// 	wantItem:  test.Item{Value: 21},
		// 	wantFound: true,
		// },
		"one-match": {
			items:     test.ItemDict,
			predicate: func(_ int, item test.Item) bool { return item.Value < 60 },
			wantKey: 40,
			wantItem:  test.Item{Value: 87},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIfNot[int, test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindIfNotStructPointer(t *testing.T) {

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
			items:     test.ItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value < 100 },
			wantKey: 0,
			wantItem:  nil,
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     test.ItemPointerDict,
		// 	predicate: func(_ int, item *test.Item) bool { return item.Value%10 == 4 },
		//  wantKey:   10,
		// 	wantItem:  &test.Item{Value: 21},
		// 	wantFound: true,
		// },
		"one-match": {
			items:     test.ItemPointerDict,
			predicate: func(_ int, item *test.Item) bool { return item.Value < 60 },
			wantKey:  40,
			wantItem:  &test.Item{Value: 87},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIfNot[int, *test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

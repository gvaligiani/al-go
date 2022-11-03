package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestFindIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, int64]
		predicate dict.Predicate[int, int64]
		wantKey   int
		wantItem  int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantKey:   0,
			wantItem:  0,
			wantFound: false,
		},
		"empty": {
			items:     EmptyInt64Dict,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantKey:   0,
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantKey:   0,
			wantItem:  0,
			wantFound: false,
		},
		"one-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i%10 == 4 },
			wantKey:   30,
			wantItem:  34,
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     DefaultInt64Dict,
		// 	predicate: func(_ int, i int64) bool { return i%10 == 2 },
		// 	wantKey:   20,
		// 	wantItem:  12,
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIf(testCase.items, testCase.predicate)

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
		items     dict.Dict[int, Item]
		predicate dict.Predicate[int, Item]
		wantKey   int
		wantItem  Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantKey:   0,
			wantItem:  Item{},
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemDict,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantKey:   0,
			wantItem:  Item{},
			wantFound: false,
		},
		"no-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantKey:   0,
			wantItem:  Item{},
			wantFound: false,
		},
		"one-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 4 },
			wantKey:   30,
			wantItem:  Item{Value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     ItemDict,
		// 	predicate: func(_ int, item Item) bool { return item.Value%10 == 2 },
		// 	wantKey:   20,
		// 	wantItem:  Item{Value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIf(testCase.items, testCase.predicate)

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
		items     dict.Dict[int, *Item]
		predicate dict.Predicate[int, *Item]
		wantKey   int
		wantItem  *Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantKey:   0,
			wantItem:  nil,
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantKey:   0,
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantKey:   0,
			wantItem:  nil,
			wantFound: false,
		},
		"one-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 4 },
			wantKey:   30,
			wantItem:  &Item{Value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     ItemPointerDict,
		// 	predicate: func(_ int, item *Item) bool { return item.Value%10 == 2 },
		// 	wantKey:   20,
		// 	wantItem:  &Item{Value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

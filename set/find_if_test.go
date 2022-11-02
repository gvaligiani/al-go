package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestFindIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int64]struct{}
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
			items:     EmptyInt64Set,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"one-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i%10 == 4 },
			wantItem:  34,
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     test.Int64Set,
		// 	predicate: func(i int64) bool { return i%10 == 2 },
		// 	wantItem:  12,
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIf[int64](testCase.items, testCase.predicate)

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
		items     map[Item]struct{}
		predicate func(Item) bool
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
			items:     EmptyItemSet,
			predicate: func(item Item) bool { return item.Value%10 == 3 },
			wantItem:  Item{},
			wantFound: false,
		},
		"no-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value%10 == 3 },
			wantItem:  Item{},
			wantFound: false,
		},
		"one-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value%10 == 4 },
			wantItem:  Item{Value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     ItemSet,
		// 	predicate: func(item Item) bool { return item.Value%10 == 2 },
		// 	wantItem:  Item{Value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIf[Item](testCase.items, testCase.predicate)

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
		items     map[*Item]struct{}
		predicate func(*Item) bool
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
			items:     EmptyItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"one-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%10 == 4 },
			wantItem:  &Item{Value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     ItemPointerSet,
		// 	predicate: func(item *Item) bool { return item.Value%10 == 2 },
		// 	wantItem:  &Item{Value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIf[*Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

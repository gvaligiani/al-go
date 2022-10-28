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
			items:     test.EmptyInt64Set,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     test.Int64Set,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"one-match": {
			items:     test.Int64Set,
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
		items     map[test.Item]struct{}
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
			items:     test.EmptyItemSet,
			predicate: func(item test.Item) bool { return item.Value%10 == 3 },
			wantItem:  test.Item{},
			wantFound: false,
		},
		"no-match": {
			items:     test.ItemSet,
			predicate: func(item test.Item) bool { return item.Value%10 == 3 },
			wantItem:  test.Item{},
			wantFound: false,
		},
		"one-match": {
			items:     test.ItemSet,
			predicate: func(item test.Item) bool { return item.Value%10 == 4 },
			wantItem:  test.Item{Value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     test.ItemSet,
		// 	predicate: func(item test.Item) bool { return item.Value%10 == 2 },
		// 	wantItem:  test.Item{Value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIf[test.Item](testCase.items, testCase.predicate)

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
		items     map[*test.Item]struct{}
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
			items:     test.EmptyItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     test.ItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"one-match": {
			items:     test.ItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value%10 == 4 },
			wantItem:  &test.Item{Value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     test.ItemPointerSet,
		// 	predicate: func(item *test.Item) bool { return item.Value%10 == 2 },
		// 	wantItem:  &test.Item{Value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIf[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

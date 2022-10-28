package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestFindIfNotInt64(t *testing.T) {

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
			wantFound: false,
			wantItem:  0,
		},
		"empty": {
			items:     test.EmptyInt64List,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"no-match": {
			items:     test.DefaultInt64List,
			predicate: func(i int64) bool { return i < 100 },
			wantFound: false,
			wantItem:  0,
		},
		"one-match": {
			items:     test.DefaultInt64List,
			predicate: func(i int64) bool { return i%10 == 4 },
			wantFound: true,
			wantItem:  21,
		},
		"two-matches": {
			items:     test.DefaultInt64List,
			predicate: func(i int64) bool { return i < 60 },
			wantFound: true,
			wantItem:  87,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIfNot[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfNotStruct(t *testing.T) {

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
			wantFound: false,
			wantItem:  test.Item{},
		},
		"empty": {
			items:     test.EmptyItemList,
			predicate: func(item test.Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  test.Item{},
		},
		"no-match": {
			items:     test.DefaultItemList,
			predicate: func(item test.Item) bool { return item.Value < 100 },
			wantFound: false,
			wantItem:  test.Item{},
		},
		"one-match": {
			items:     test.DefaultItemList,
			predicate: func(item test.Item) bool { return item.Value%10 == 4 },
			wantFound: true,
			wantItem:  test.Item{Value: 21},
		},
		"two-matches": {
			items:     test.DefaultItemList,
			predicate: func(item test.Item) bool { return item.Value < 60 },
			wantFound: true,
			wantItem:  test.Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIfNot[test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfNotStructPointer(t *testing.T) {

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
			wantFound: false,
			wantItem:  nil,
		},
		"empty": {
			items:     test.EmptyItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"no-match": {
			items:     test.DefaultItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value < 100 },
			wantFound: false,
			wantItem:  nil,
		},
		"one-match": {
			items:     test.DefaultItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value%10 == 4 },
			wantFound: true,
			wantItem:  &test.Item{Value: 21},
		},
		"two-matches": {
			items:     test.DefaultItemPointerList,
			predicate: func(item *test.Item) bool { return item.Value < 60 },
			wantFound: true,
			wantItem:  &test.Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIfNot[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

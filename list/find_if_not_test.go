package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestFindIfNotInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		predicate list.Predicate[int64]
		wantIndex int
		wantItem  int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  0,
		},
		"empty": {
			items:     EmptyInt64List,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  0,
		},
		"no-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  0,
		},
		"one-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%10 == 4 },
			wantIndex: 0,
			wantFound: true,
			wantItem:  21,
		},
		"two-matches": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i < 60 },
			wantIndex: 3,
			wantFound: true,
			wantItem:  87,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotItem, gotFound := list.FindIfNot(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfNotStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[Item]
		predicate list.Predicate[Item]
		wantIndex int
		wantItem  Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  Item{},
		},
		"empty": {
			items:     EmptyItemList,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  Item{},
		},
		"no-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value < 100 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  Item{},
		},
		"one-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 4 },
			wantIndex: 0,
			wantFound: true,
			wantItem:  Item{Value: 21},
		},
		"two-matches": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value < 60 },
			wantIndex: 3,
			wantFound: true,
			wantItem:  Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotItem, gotFound := list.FindIfNot(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfNotStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[*Item]
		predicate list.Predicate[*Item]
		wantIndex int
		wantItem  *Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  nil,
		},
		"empty": {
			items:     EmptyItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  nil,
		},
		"no-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value < 100 },
			wantIndex: -1,
			wantFound: false,
			wantItem:  nil,
		},
		"one-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 4 },
			wantIndex: 0,
			wantFound: true,
			wantItem:  &Item{Value: 21},
		},
		"two-matches": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value < 60 },
			wantIndex: 3,
			wantFound: true,
			wantItem:  &Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotItem, gotFound := list.FindIfNot(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

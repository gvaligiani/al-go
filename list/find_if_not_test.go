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
		wantItem  int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"empty": {
			items:     EmptyInt64List,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"no-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantFound: false,
			wantItem:  0,
		},
		"one-match": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i%10 == 4 },
			wantFound: true,
			wantItem:  21,
		},
		"two-matches": {
			items:     DefaultInt64List,
			predicate: func(_ int, i int64) bool { return i < 60 },
			wantFound: true,
			wantItem:  87,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

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
		items     list.List[Item]
		predicate list.Predicate[Item]
		wantItem  Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  Item{},
		},
		"empty": {
			items:     EmptyItemList,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  Item{},
		},
		"no-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value < 100 },
			wantFound: false,
			wantItem:  Item{},
		},
		"one-match": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value%10 == 4 },
			wantFound: true,
			wantItem:  Item{Value: 21},
		},
		"two-matches": {
			items:     DefaultItemList,
			predicate: func(_ int, item Item) bool { return item.Value < 60 },
			wantFound: true,
			wantItem:  Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIfNot[Item](testCase.items, testCase.predicate)

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
		items     list.List[*Item]
		predicate list.Predicate[*Item]
		wantItem  *Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"empty": {
			items:     EmptyItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"no-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value < 100 },
			wantFound: false,
			wantItem:  nil,
		},
		"one-match": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value%10 == 4 },
			wantFound: true,
			wantItem:  &Item{Value: 21},
		},
		"two-matches": {
			items:     DefaultItemPointerList,
			predicate: func(_ int, item *Item) bool { return item.Value < 60 },
			wantFound: true,
			wantItem:  &Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := list.FindIfNot[*Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

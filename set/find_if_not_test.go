package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
	"github.com/gvaligiani/al.go/util"
)

func TestFindIfNotInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[int64]
		predicate util.Predicate[int64]
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
			items:     EmptyInt64Set,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"no-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i < 100 },
			wantFound: false,
			wantItem:  0,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     test.Int64Set,
		// 	predicate: func(i int64) bool { return i%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  21,
		// },
		"one-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i < 60 },
			wantFound: true,
			wantItem:  87,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot(testCase.items, testCase.predicate)

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
		items     set.Set[Item]
		predicate func(Item) bool
		wantItem  Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  Item{},
		},
		"empty": {
			items:     EmptyItemSet,
			predicate: func(item Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  Item{},
		},
		"no-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value < 100 },
			wantFound: false,
			wantItem:  Item{},
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     ItemSet,
		// 	predicate: func(item Item) bool { return item.Value%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  Item{Value: 21},
		// },
		"one-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value < 60 },
			wantFound: true,
			wantItem:  Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot(testCase.items, testCase.predicate)

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
		items     set.Set[*Item]
		predicate func(*Item) bool
		wantItem  *Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"empty": {
			items:     EmptyItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"no-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value < 100 },
			wantFound: false,
			wantItem:  nil,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     ItemPointerSet,
		// 	predicate: func(item *Item) bool { return item.Value%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  &Item{Value: 21},
		// },
		"one-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value < 60 },
			wantFound: true,
			wantItem:  &Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestFindIfNotInt64(t *testing.T) {

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
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantKey:   0,
			wantItem:  0,
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     DefaultInt64Dict,
		// 	predicate: func(_ int, i int64) bool { return i%10 == 4 },
		// 	wantKey:   10,
		// 	wantItem:  21,
		// 	wantFound: true,
		// },
		"one-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i < 60 },
			wantKey:   40,
			wantItem:  87,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIfNot(testCase.items, testCase.predicate)

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
			predicate: func(_ int, item Item) bool { return item.Value < 100 },
			wantKey:   0,
			wantItem:  Item{},
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     ItemDict,
		// 	predicate: func(_ int, item Item) bool { return item.Value%10 == 4 },
		// 	wantKey:   10,
		// 	wantItem:  Item{Value: 21},
		// 	wantFound: true,
		// },
		"one-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value < 60 },
			wantKey:   40,
			wantItem:  Item{Value: 87},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIfNot(testCase.items, testCase.predicate)

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
			predicate: func(_ int, item *Item) bool { return item.Value < 100 },
			wantKey:   0,
			wantItem:  nil,
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     ItemPointerDict,
		// 	predicate: func(_ int, item *Item) bool { return item.Value%10 == 4 },
		//  wantKey:   10,
		// 	wantItem:  &Item{Value: 21},
		// 	wantFound: true,
		// },
		"one-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value < 60 },
			wantKey:   40,
			wantItem:  &Item{Value: 87},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIfNot(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

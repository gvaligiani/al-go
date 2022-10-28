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

	empty := map[int]int64{}
	items := map[int]int64{
		10: 21,
		20: 12,
		30: 34,
		40: 87,
		50: 52,
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
			items:     empty,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantKey: 0,
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantKey: 0,
			wantItem:  0,
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(_ int, i int64) bool { return i%10 == 4 },
		// 	wantKey:   10,
		// 	wantItem:  21,
		// 	wantFound: true,
		// },
		"one-match": {
			items:     items,
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

	type Item struct {
		value int
	}

	type TestCase struct {
		items     map[int]Item
		predicate func(int, Item) bool
		wantKey int
		wantItem  Item
		wantFound bool
	}

	empty := map[int]Item{}
	items := map[int]Item{
		10: {value: 21},
		20: {value: 12},
		30: {value: 34},
		40: {value: 87},
		50: {value: 52},
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i Item) bool { return i.value%10 == 3 },
			wantKey: 0,
			wantItem:  Item{},
			wantFound: false,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i Item) bool { return i.value%10 == 3 },
			wantKey: 0,
			wantItem:  Item{},
			wantFound: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value < 100 },
			wantKey: 0,
			wantItem:  Item{},
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(_ int, i Item) bool { return i.value%10 == 4 },
		// 	wantKey:   10,
		// 	wantItem:  Item{value: 21},
		// 	wantFound: true,
		// },
		"one-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value < 60 },
			wantKey: 40,
			wantItem:  Item{value: 87},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIfNot[int, Item](testCase.items, testCase.predicate)

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

	type Item struct {
		value int
	}

	type TestCase struct {
		items     map[int]*Item
		predicate func(int, *Item) bool
		wantKey int
		wantItem  *Item
		wantFound bool
	}

	empty := map[int]*Item{}
	items := map[int]*Item{
		10: {value: 21},
		20: {value: 12},
		30: {value: 34},
		40: {value: 87},
		50: {value: 52},
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i *Item) bool { return i.value%10 == 3 },
			wantKey: 0,
			wantItem:  nil,
			wantFound: false,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i *Item) bool { return i.value%10 == 3 },
			wantKey: 0,
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value < 100 },
			wantKey: 0,
			wantItem:  nil,
			wantFound: false,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(_ int, i *Item) bool { return i.value%10 == 4 },
		//  wantKey:   10,
		// 	wantItem:  &Item{value: 21},
		// 	wantFound: true,
		// },
		"one-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value < 60 },
			wantKey:  40,
			wantItem:  &Item{value: 87},
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotKey, gotItem, gotFound := dict.FindIfNot[int, *Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantKey, gotKey, "wrong key!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

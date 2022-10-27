package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestFindIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		predicate func(int, int64) bool
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
			wantItem:  0,
			wantFound: false,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantItem:  0,
			wantFound: false,
		},
		"one-match": {
			items:     items,
			predicate: func(_ int, i int64) bool { return i%10 == 4 },
			wantItem:  34,
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     items,
		// 	predicate: func(_ int, i int64) bool { return i%10 == 2 },
		// 	wantItem:  12,
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := dict.FindIf[int,int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items     map[int]Item
		predicate func(int, Item) bool
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
			wantItem:  Item{},
			wantFound: false,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i Item) bool { return i.value%10 == 3 },
			wantItem:  Item{},
			wantFound: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value%10 == 3 },
			wantItem:  Item{},
			wantFound: false,
		},
		"one-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value%10 == 4 },
			wantItem:  Item{value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     items,
		// 	predicate: func(_ int, i Item) bool { return i.value%10 == 2 },
		// 	wantItem:  Item{value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := dict.FindIf[int, Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items     map[int]*Item
		predicate func(int, *Item) bool
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
			wantItem:  nil,
			wantFound: false,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i *Item) bool { return i.value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value%10 == 3 },
			wantItem:  nil,
			wantFound: false,
		},
		"one-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value%10 == 4 },
			wantItem:  &Item{value: 34},
			wantFound: true,
		},
		// NOTE order could not be guarantee >>> expected 12 or 52
		// "two-matches": {
		// 	items:     items,
		// 	predicate: func(_ int, i *Item) bool { return i.value%10 == 2 },
		// 	wantItem:  &Item{value: 12},
		// 	wantFound: true,
		// },
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := dict.FindIf[int, *Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

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
			wantFound: false,
			wantItem:  0,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantFound: false,
			wantItem:  0,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(_ int, i int64) bool { return i%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  21,
		// },
		"one-match": {
			items:     items,
			predicate: func(_ int, i int64) bool { return i < 50 },
			wantFound: true,
			wantItem:  87,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := dict.FindIfNot[int, int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
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
			wantFound: false,
			wantItem:  Item{},
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i Item) bool { return i.value%10 == 3 },
			wantFound: false,
			wantItem:  Item{},
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value < 100 },
			wantFound: false,
			wantItem:  Item{},
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(_ int, i Item) bool { return i.value%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  Item{value: 21},
		// },
		"one-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value < 50 },
			wantFound: true,
			wantItem:  Item{value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := dict.FindIfNot[int, Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
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
			wantFound: false,
			wantItem:  nil,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i *Item) bool { return i.value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value < 100 },
			wantFound: false,
			wantItem:  nil,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(_ int, i *Item) bool { return i.value%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  &Item{value: 21},
		// },
		"one-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value < 50 },
			wantFound: true,
			wantItem:  &Item{value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := dict.FindIfNot[int, *Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

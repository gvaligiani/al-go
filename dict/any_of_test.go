package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestAnyOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		predicate func(int, int64) bool
		wantAnyOf bool
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
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(_ int, i int64) bool { return i > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     items,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf[int,int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items     map[int]Item
		predicate func(int, Item) bool
		wantAnyOf bool
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
			predicate: func(_ int, i Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     items,
			predicate: func(_ int, i Item) bool { return i.value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf[int, Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items     map[int]*Item
		predicate func(int, *Item) bool
		wantAnyOf bool
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
			predicate: func(_ int, i *Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     empty,
			predicate: func(_ int, i *Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     items,
			predicate: func(_ int, i *Item) bool { return i.value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf[int, *Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

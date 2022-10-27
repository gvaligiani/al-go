package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestNoneOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      map[int]int64
		predicate  func(int, int64) bool
		wantNoneOf bool
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
			items:      nil,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      empty,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      items,
			predicate:  func(_ int, i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      items,
			predicate:  func(_ int, i int64) bool { return i > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      items,
			predicate:  func(_ int, i int64) bool { return i < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := dict.NoneOf[int,int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items      map[int]Item
		predicate  func(int, Item) bool
		wantNoneOf bool
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
			items:      nil,
			predicate:  func(_ int, i Item) bool { return i.value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      empty,
			predicate:  func(_ int, i Item) bool { return i.value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      items,
			predicate:  func(_ int, i Item) bool { return i.value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      items,
			predicate:  func(_ int, i Item) bool { return i.value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      items,
			predicate:  func(_ int, i Item) bool { return i.value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := dict.NoneOf[int,Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items      map[int]*Item
		predicate  func(int, *Item) bool
		wantNoneOf bool
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
			items:      nil,
			predicate:  func(_ int, i *Item) bool { return i.value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      empty,
			predicate:  func(_ int, i *Item) bool { return i.value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      items,
			predicate:  func(_ int, i *Item) bool { return i.value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      items,
			predicate:  func(_ int, i *Item) bool { return i.value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      items,
			predicate:  func(_ int, i *Item) bool { return i.value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := dict.NoneOf[int,*Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

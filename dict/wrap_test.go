package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestWrapInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items  map[int]int64
		wantSum int64
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
			items:  nil,
			wantSum: 0,
		},
		"empty": {
			items:  empty,
			wantSum: 0,
		},
		"all": {
			items:  items,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		dict.Wrap[int,int64](testCase.items).Each(func(_ int, item int64) { gotSum += item })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

func TestWrapStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items  map[int]Item
		wantSum int64
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
			items:  nil,
			wantSum: 0,
		},
		"empty": {
			items:  empty,
			wantSum: 0,
		},
		"all": {
			items:  items,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		dict.Wrap[int, Item](testCase.items).Each(func(_ int, item Item) { gotSum += item.value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

func TestWrapStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items  map[int]*Item
		wantSum int64
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
			items:  nil,
			wantSum: 0,
		},
		"empty": {
			items:  empty,
			wantSum: 0,
		},
		"all": {
			items:  items,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		dict.Wrap[int,*Item](testCase.items).Each(func(_ int, item *Item) { gotSum += item.value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

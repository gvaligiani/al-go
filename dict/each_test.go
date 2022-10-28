package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestEachInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items  map[int]int64
		wantNb int
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
			wantNb: 0,
		},
		"empty": {
			items:  empty,
			wantNb: 0,
		},
		"all": {
			items:  items,
			wantNb: 5,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		dict.Each[int,int64](testCase.items, func(_ int, _ int64) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

func TestEachStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items  map[int]Item
		wantNb int
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
			wantNb: 0,
		},
		"empty": {
			items:  empty,
			wantNb: 0,
		},
		"all": {
			items:  items,
			wantNb: 5,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		dict.Each[int,Item](testCase.items, func(_ int, _ Item) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

func TestEachStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items  map[int]*Item
		wantNb int
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
			wantNb: 0,
		},
		"empty": {
			items:  empty,
			wantNb: 0,
		},
		"all": {
			items:  items,
			wantNb: 5,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		dict.Each[int, *Item](testCase.items, func(_ int, _ *Item) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

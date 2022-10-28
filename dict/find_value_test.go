package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestFindValueInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		value     int64
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
			value:     34,
			wantFound: false,
		},
		"empty": {
			items:     empty,
			value:     34,
			wantFound: false,
		},
		"found": {
			items:     items,
			value:     34,
			wantFound: true,
		},
		"not-found": {
			items:     items,
			value:     99,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindValue[int,int64](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindValueStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items     map[int]Item
		value     Item
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
			value:     Item{value: 34},
			wantFound: false,
		},
		"empty": {
			items:     empty,
			value:     Item{value: 34},
			wantFound: false,
		},
		"found": {
			items:     items,
			value:     Item{value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     items,
			value:     Item{value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindValue[int,Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindValueStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items     map[int]*Item
		value     *Item
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
			value:     &Item{value: 34},
			wantFound: false,
		},
		"empty": {
			items:     empty,
			value:     &Item{value: 34},
			wantFound: false,
		},
		"found": {
			items:     items,
			value:     &Item{value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     items,
			value:     &Item{value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindValue[int, *Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

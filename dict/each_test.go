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

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantNb: 0,
		},
		"empty": {
			items:  test.EmptyInt64Dict,
			wantNb: 0,
		},
		"all": {
			items:  test.Int64Dict,
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

	type TestCase struct {
		items  map[int]test.Item
		wantNb int
	}

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantNb: 0,
		},
		"empty": {
			items:  test.EmptyItemDict,
			wantNb: 0,
		},
		"all": {
			items:  test.ItemDict,
			wantNb: 5,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		dict.Each[int,test.Item](testCase.items, func(_ int, _ test.Item) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

func TestEachStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items  map[int]*test.Item
		wantNb int
	}

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantNb: 0,
		},
		"empty": {
			items:  test.EmptyItemPointerDict,
			wantNb: 0,
		},
		"all": {
			items:  test.ItemPointerDict,
			wantNb: 5,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		dict.Each[int, *test.Item](testCase.items, func(_ int, _ *test.Item) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

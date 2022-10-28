package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestEachInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items  []int64
		wantNb int
	}

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantNb: 0,
		},
		"empty": {
			items:  test.EmptyInt64List,
			wantNb: 0,
		},
		"all": {
			items:  test.Int64List,
			wantNb: 5,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		list.Each[int64](testCase.items, func(_ int, _ int64) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

func TestEachStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items  []test.Item
		wantNb int
	}

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantNb: 0,
		},
		"empty": {
			items:  test.EmptyItemList,
			wantNb: 0,
		},
		"all": {
			items:  test.ItemList,
			wantNb: 5,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		list.Each[test.Item](testCase.items, func(_ int, _ test.Item) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

func TestEachStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items  []*test.Item
		wantNb int
	}

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantNb: 0,
		},
		"empty": {
			items:  test.EmptyItemPointerList,
			wantNb: 0,
		},
		"all": {
			items:  test.ItemPointerList,
			wantNb: 5,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		list.Each[*test.Item](testCase.items, func(_ int, _ *test.Item) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

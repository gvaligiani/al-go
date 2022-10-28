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

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantSum: 0,
		},
		"empty": {
			items:  test.EmptyInt64Dict,
			wantSum: 0,
		},
		"all": {
			items:  test.DefaultInt64Dict,
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

	type TestCase struct {
		items  map[int]test.Item
		wantSum int64
	}

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantSum: 0,
		},
		"empty": {
			items:  test.EmptyItemDict,
			wantSum: 0,
		},
		"all": {
			items:  test.DefaultItemDict,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		dict.Wrap[int, test.Item](testCase.items).Each(func(_ int, item test.Item) { gotSum += item.Value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

func TestWrapStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items  map[int]*test.Item
		wantSum int64
	}

	testCases := map[string]TestCase{
		"nil": {
			items:  nil,
			wantSum: 0,
		},
		"empty": {
			items:  test.EmptyItemPointerDict,
			wantSum: 0,
		},
		"all": {
			items:  test.DefaultItemPointerDict,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		dict.Wrap[int,*test.Item](testCase.items).Each(func(_ int, item *test.Item) { gotSum += item.Value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

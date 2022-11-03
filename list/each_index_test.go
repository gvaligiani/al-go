package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestEachIndexInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items   []int64
		wantSum int64
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   EmptyInt64List,
			wantSum: 0,
		},
		"all": {
			items:   DefaultInt64List,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		list.EachIndex[int64](testCase.items, func(_ int, item int64) { gotSum += item })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

func TestEachIndexStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items   []Item
		wantSum int64
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   EmptyItemList,
			wantSum: 0,
		},
		"all": {
			items:   DefaultItemList,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		list.EachIndex[Item](testCase.items, func(_ int, item Item) { gotSum += item.Value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

func TestEachIndexStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items   []*Item
		wantSum int64
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   EmptyItemPointerList,
			wantSum: 0,
		},
		"all": {
			items:   DefaultItemPointerList,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		list.EachIndex[*Item](testCase.items, func(_ int, item *Item) { gotSum += item.Value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

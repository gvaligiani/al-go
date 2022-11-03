package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
)

func TestEachInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items   set.Set[int64]
		wantSum int64
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   EmptyInt64Set,
			wantSum: 0,
		},
		"all": {
			items:   DefaultInt64Set,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		set.Each(testCase.items, func(item int64) { gotSum += item })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

func TestEachStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items   set.Set[Item]
		wantSum int64
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   EmptyItemSet,
			wantSum: 0,
		},
		"all": {
			items:   DefaultItemSet,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		set.Each(testCase.items, func(item Item) { gotSum += item.Value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

func TestEachStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items   set.Set[*Item]
		wantSum int64
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   EmptyItemPointerSet,
			wantSum: 0,
		},
		"all": {
			items:   DefaultItemPointerSet,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		set.Each(testCase.items, func(item *Item) { gotSum += item.Value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestWrapInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items   map[int64]struct{}
		wantSum int64
	}

	empty := map[int64]struct{}{}
	items := map[int64]struct{}{
		21: {},
		12: {},
		34: {},
		87: {},
		52: {},
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   empty,
			wantSum: 0,
		},
		"all": {
			items:   items,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		set.Wrap[int64](testCase.items).Each(func(item int64) { gotSum += item })

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
		items   map[Item]struct{}
		wantSum int64
	}

	empty := map[Item]struct{}{}
	items := map[Item]struct{}{
		{value: 21}: {},
		{value: 12}: {},
		{value: 34}: {},
		{value: 87}: {},
		{value: 52}: {},
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   empty,
			wantSum: 0,
		},
		"all": {
			items:   items,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		set.Wrap[Item](testCase.items).Each(func(item Item) { gotSum += item.value })

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
		items   map[*Item]struct{}
		wantSum int64
	}

	empty := map[*Item]struct{}{}
	items := map[*Item]struct{}{
		{value: 21}: {},
		{value: 12}: {},
		{value: 34}: {},
		{value: 87}: {},
		{value: 52}: {},
	}

	testCases := map[string]TestCase{
		"nil": {
			items:   nil,
			wantSum: 0,
		},
		"empty": {
			items:   empty,
			wantSum: 0,
		},
		"all": {
			items:   items,
			wantSum: 21 + 12 + 34 + 87 + 52,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		var gotSum int64
		set.Wrap[*Item](testCase.items).Each(func(item *Item) { gotSum += item.value })

		// assert
		require.Equalf(t, testCase.wantSum, gotSum, "wrong sum!")
	})
}

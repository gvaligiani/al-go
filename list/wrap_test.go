package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestWrapInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items  []int64
		wantNb int
	}

	empty := []int64{}
	items := []int64{
		21,
		12,
		34,
		87,
		52,
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

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		list.Wrap[int64](testCase.items).Each(func(_ int, _ int64) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

func TestWrapStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items  []Item
		wantNb int
	}

	empty := []Item{}
	items := []Item{
		{value: 21},
		{value: 12},
		{value: 34},
		{value: 87},
		{value: 52},
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

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		list.Wrap[Item](testCase.items).Each(func(_ int, _ Item) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

func TestWrapStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items  []*Item
		wantNb int
	}

	empty := []*Item{}
	items := []*Item{
		{value: 21},
		{value: 12},
		{value: 34},
		{value: 87},
		{value: 52},
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

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNb := 0
		list.Wrap[*Item](testCase.items).Each(func(_ int, _ *Item) { gotNb++ })

		// assert
		require.Equalf(t, testCase.wantNb, gotNb, "wrong nb!")
	})
}

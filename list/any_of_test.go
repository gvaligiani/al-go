package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestAnyOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		predicate func(int64) bool
		wantAnyOf bool
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
			items:     nil,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     empty,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     items,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(i int64) bool { return i > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     items,
			predicate: func(i int64) bool { return i < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items     []Item
		predicate func(Item) bool
		wantAnyOf bool
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
			items:     nil,
			predicate: func(i Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     empty,
			predicate: func(i Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     items,
			predicate: func(i Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(i Item) bool { return i.value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     items,
			predicate: func(i Item) bool { return i.value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf[Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int
	}

	type TestCase struct {
		items     []*Item
		predicate func(*Item) bool
		wantAnyOf bool
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
			items:     nil,
			predicate: func(i *Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     empty,
			predicate: func(i *Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     items,
			predicate: func(i *Item) bool { return i.value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(i *Item) bool { return i.value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     items,
			predicate: func(i *Item) bool { return i.value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(_ string, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := list.AnyOf[*Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

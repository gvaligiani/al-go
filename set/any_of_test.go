package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestAnyOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int64]struct{}
		predicate func(int64) bool
		wantAnyOf bool
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

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := set.AnyOf[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items     map[Item]struct{}
		predicate func(Item) bool
		wantAnyOf bool
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

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := set.AnyOf[Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items     map[*Item]struct{}
		predicate func(*Item) bool
		wantAnyOf bool
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

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := set.AnyOf[*Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

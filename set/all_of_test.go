package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestAllOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int64]struct{}
		predicate func(int64) bool
		wantAllOf bool
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
			wantAllOf: true,
		},
		"empty": {
			items:     empty,
			predicate: func(i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     items,
			predicate: func(i int64) bool { return i > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(i int64) bool { return i > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     items,
			predicate: func(i int64) bool { return i < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := set.AllOf[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items     map[Item]struct{}
		predicate func(Item) bool
		wantAllOf bool
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
			wantAllOf: true,
		},
		"empty": {
			items:     empty,
			predicate: func(i Item) bool { return i.value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     items,
			predicate: func(i Item) bool { return i.value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(i Item) bool { return i.value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     items,
			predicate: func(i Item) bool { return i.value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := set.AllOf[Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items     map[*Item]struct{}
		predicate func(*Item) bool
		wantAllOf bool
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
			wantAllOf: true,
		},
		"empty": {
			items:     empty,
			predicate: func(i *Item) bool { return i.value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     items,
			predicate: func(i *Item) bool { return i.value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     items,
			predicate: func(i *Item) bool { return i.value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     items,
			predicate: func(i *Item) bool { return i.value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := set.AllOf[*Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

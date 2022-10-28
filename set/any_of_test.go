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

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyInt64Set,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.Int64Set,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.Int64Set,
			predicate: func(i int64) bool { return i > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.Int64Set,
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

	type TestCase struct {
		items     map[test.Item]struct{}
		predicate func(test.Item) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyItemSet,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.ItemSet,
			predicate: func(item test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.ItemSet,
			predicate: func(item test.Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.ItemSet,
			predicate: func(item test.Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := set.AnyOf[test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[*test.Item]struct{}
		predicate func(*test.Item) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     test.EmptyItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     test.ItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     test.ItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     test.ItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := set.AnyOf[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

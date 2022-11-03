package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
)

func TestAnyOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[int64]
		predicate set.Predicate[int64]
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyInt64Set,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := set.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[Item]
		predicate func(Item) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyItemSet,
			predicate: func(item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := set.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[*Item]
		predicate func(*Item) bool
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyItemPointerSet,
			predicate: func(item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := set.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

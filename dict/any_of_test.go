package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestAnyOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, int64]
		predicate dict.Predicate[int, int64]
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, Item]
		predicate dict.Predicate[int, Item]
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyItemDict,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

func TestAnyOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, *Item]
		predicate dict.Predicate[int, *Item]
		wantAnyOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"empty": {
			items:     EmptyItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"no-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAnyOf: false,
		},
		"some-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value > 20 },
			wantAnyOf: true,
		},
		"all-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value < 100 },
			wantAnyOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAnyOf := dict.AnyOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAnyOf, gotAnyOf, "wrong any_of!")
	})
}

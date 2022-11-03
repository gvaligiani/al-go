package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestAllOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, int64]
		predicate dict.Predicate[int, int64]
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultInt64Dict,
			predicate: func(_ int, i int64) bool { return i < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := dict.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, Item]
		predicate dict.Predicate[int, Item]
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyItemDict,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultItemDict,
			predicate: func(_ int, item Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := dict.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, *Item]
		predicate dict.Predicate[int, *Item]
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultItemPointerDict,
			predicate: func(_ int, item *Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := dict.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

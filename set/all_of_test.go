package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
)

func TestAllOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[int64]
		predicate set.Predicate[int64]
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyInt64Set,
			predicate: func(i int64) bool { return i > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := set.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[Item]
		predicate func(Item) bool
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyItemSet,
			predicate: func(item Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := set.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

func TestAllOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[*Item]
		predicate func(*Item) bool
		wantAllOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"empty": {
			items:     EmptyItemPointerSet,
			predicate: func(item *Item) bool { return item.Value > 100 },
			wantAllOf: true,
		},
		"no-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value > 100 },
			wantAllOf: false,
		},
		"some-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value > 20 },
			wantAllOf: false,
		},
		"all-match": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value < 100 },
			wantAllOf: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotAllOf := set.AllOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantAllOf, gotAllOf, "wrong all_of!")
	})
}

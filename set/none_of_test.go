package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
	"github.com/gvaligiani/al.go/util"
)

func TestNoneOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      set.Set[int64]
		predicate  util.Predicate[int64]
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      EmptyInt64Set,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      DefaultInt64Set,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      DefaultInt64Set,
			predicate:  func(i int64) bool { return i > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      DefaultInt64Set,
			predicate:  func(i int64) bool { return i < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := set.NoneOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      set.Set[Item]
		predicate  func(Item) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(item Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      EmptyItemSet,
			predicate:  func(item Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      DefaultItemSet,
			predicate:  func(item Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      DefaultItemSet,
			predicate:  func(item Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      DefaultItemSet,
			predicate:  func(item Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := set.NoneOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      set.Set[*Item]
		predicate  func(*Item) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(item *Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      EmptyItemPointerSet,
			predicate:  func(item *Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      DefaultItemPointerSet,
			predicate:  func(item *Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      DefaultItemPointerSet,
			predicate:  func(item *Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      DefaultItemPointerSet,
			predicate:  func(item *Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := set.NoneOf(testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

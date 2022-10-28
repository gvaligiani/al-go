package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestNoneOfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      map[int64]struct{}
		predicate  func(int64) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyInt64Set,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.DefaultInt64Set,
			predicate:  func(i int64) bool { return i > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.DefaultInt64Set,
			predicate:  func(i int64) bool { return i > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.DefaultInt64Set,
			predicate:  func(i int64) bool { return i < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := set.NoneOf[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      map[test.Item]struct{}
		predicate  func(test.Item) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyItemSet,
			predicate:  func(item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.DefaultItemSet,
			predicate:  func(item test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.DefaultItemSet,
			predicate:  func(item test.Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.DefaultItemSet,
			predicate:  func(item test.Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := set.NoneOf[test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

func TestNoneOfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items      map[*test.Item]struct{}
		predicate  func(*test.Item) bool
		wantNoneOf bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:      nil,
			predicate:  func(item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"empty": {
			items:      test.EmptyItemPointerSet,
			predicate:  func(item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"no-match": {
			items:      test.DefaultItemPointerSet,
			predicate:  func(item *test.Item) bool { return item.Value > 100 },
			wantNoneOf: true,
		},
		"some-match": {
			items:      test.DefaultItemPointerSet,
			predicate:  func(item *test.Item) bool { return item.Value > 20 },
			wantNoneOf: false,
		},
		"all-match": {
			items:      test.DefaultItemPointerSet,
			predicate:  func(item *test.Item) bool { return item.Value < 100 },
			wantNoneOf: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotNoneOf := set.NoneOf[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantNoneOf, gotNoneOf, "wrong none_of!")
	})
}

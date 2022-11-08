package set_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
	"github.com/stretchr/testify/require"
)

func TestKeepIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       set.Set[int64]
		predicate   set.Predicate[int64]
		wantUpdated bool
		wantItems   set.Set[int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(i int64) bool { return i%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyInt64Set,
			predicate:   func(i int64) bool { return i%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyInt64Set,
		},
		"keep-none": {
			items:       DefaultInt64Set,
			predicate:   func(i int64) bool { return false },
			wantUpdated: true,
			wantItems:   EmptyInt64Set,
		},
		"keep-odd": {
			items:       DefaultInt64Set,
			predicate:   func(i int64) bool { return i%2 == 0 },
			wantUpdated: true,
			wantItems: set.Set[int64]{
				12: {},
				34: {},
				52: {},
			},
		},
		"keep-even": {
			items:       DefaultInt64Set,
			predicate:   func(i int64) bool { return i%2 == 1 },
			wantUpdated: true,
			wantItems: set.Set[int64]{
				21: {},
				87: {},
			},
		},
		"keep-all": {
			items:       DefaultInt64Set,
			predicate:   func(i int64) bool { return true },
			wantUpdated: false,
			wantItems:   DefaultInt64Set,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)
		gotUpdated := set.KeepIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestKeepIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       set.Set[Item]
		predicate   set.Predicate[Item]
		wantUpdated bool
		wantItems   set.Set[Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(item Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyItemSet,
			predicate:   func(item Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyItemSet,
		},
		"keep-none": {
			items:       DefaultItemSet,
			predicate:   func(item Item) bool { return false },
			wantUpdated: true,
			wantItems:   EmptyItemSet,
		},
		"keep-odd": {
			items:       DefaultItemSet,
			predicate:   func(item Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: set.Set[Item]{
				Item{Value: 12}: {},
				Item{Value: 34}: {},
				Item{Value: 52}: {},
			},
		},
		"keep-even": {
			items:       DefaultItemSet,
			predicate:   func(item Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: set.Set[Item]{
				Item{Value: 21}: {},
				Item{Value: 87}: {},
			},
		},
		"keep-all": {
			items:       DefaultItemSet,
			predicate:   func(item Item) bool { return true },
			wantUpdated: false,
			wantItems:   DefaultItemSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)
		gotUpdated := set.KeepIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestKeepIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       set.Set[*Item]
		predicate   set.Predicate[*Item]
		wantUpdated bool
		wantItems   set.Set[*Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyItemPointerSet,
			predicate:   func(item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyItemPointerSet,
		},
		"keep-none": {
			items:       DefaultItemPointerSet,
			predicate:   func(item *Item) bool { return false },
			wantUpdated: true,
			wantItems:   EmptyItemPointerSet,
		},
		"keep-odd": {
			items:       DefaultItemPointerSet,
			predicate:   func(item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: set.Set[*Item]{
				&Item{Value: 12}: {},
				&Item{Value: 34}: {},
				&Item{Value: 52}: {},
			},
		},
		"keep-even": {
			items:       DefaultItemPointerSet,
			predicate:   func(item *Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: set.Set[*Item]{
				&Item{Value: 21}: {},
				&Item{Value: 87}: {},
			},
		},
		"keep-all": {
			items:       DefaultItemPointerSet,
			predicate:   func(item *Item) bool { return true },
			wantUpdated: false,
			wantItems:   DefaultItemPointerSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)
		gotUpdated := set.KeepIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

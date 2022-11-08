package set_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
	"github.com/gvaligiani/al.go/util"
	"github.com/stretchr/testify/require"
)

func TestRemoveIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       set.Set[int64]
		predicate   util.Predicate[int64]
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
		"remove-none": {
			items:       DefaultInt64Set,
			predicate:   func(i int64) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultInt64Set,
		},
		"remove-odd": {
			items:       DefaultInt64Set,
			predicate:   func(i int64) bool { return i%2 == 0 },
			wantUpdated: true,
			wantItems: set.New[int64](
				21,
				87,
			),
		},
		"remove-even": {
			items:       DefaultInt64Set,
			predicate:   func(i int64) bool { return i%2 == 1 },
			wantUpdated: true,
			wantItems: set.New[int64](
				12,
				34,
				52,
			),
		},
		"remove-all": {
			items:       DefaultInt64Set,
			predicate:   func(i int64) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyInt64Set,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)
		gotUpdated := set.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestRemoveIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       set.Set[Item]
		predicate   util.Predicate[Item]
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
		"remove-none": {
			items:       DefaultItemSet,
			predicate:   func(item Item) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultItemSet,
		},
		"remove-odd": {
			items:       DefaultItemSet,
			predicate:   func(item Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: set.Set[Item]{
				{Value: 21}: {},
				{Value: 87}: {},
			},
		},
		"remove-even": {
			items:       DefaultItemSet,
			predicate:   func(item Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: set.Set[Item]{
				{Value: 12}: {},
				{Value: 34}: {},
				{Value: 52}: {},
			},
		},
		"remove-all": {
			items:       DefaultItemSet,
			predicate:   func(item Item) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyItemSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)
		gotUpdated := set.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestRemoveIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       set.Set[*Item]
		predicate   util.Predicate[*Item]
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
		"remove-none": {
			items:       DefaultItemPointerSet,
			predicate:   func(item *Item) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultItemPointerSet,
		},
		"remove-odd": {
			items:       DefaultItemPointerSet,
			predicate:   func(item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: set.New(
				&Item{Value: 21},
				&Item{Value: 87},
			),
		},
		"remove-even": {
			items:       DefaultItemPointerSet,
			predicate:   func(item *Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: set.New(
				&Item{Value: 12},
				&Item{Value: 34},
				&Item{Value: 52},
			),
		},
		"remove-all": {
			items:       DefaultItemPointerSet,
			predicate:   func(item *Item) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyItemPointerSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)
		gotUpdated := set.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

package set_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
	"github.com/stretchr/testify/require"
)

func TestMergeInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left        set.Set[int64]
		right       set.Set[int64]
		wantUpdated bool
		wantItems   set.Set[int64]
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:        nil,
			right:       nil,
			wantUpdated: false,
			wantItems:   nil,
		},
		"nil-empty": {
			left:        nil,
			right:       EmptyInt64Set,
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty-nil": {
			left:        EmptyInt64Set,
			right:       nil,
			wantUpdated: false,
			wantItems:   EmptyInt64Set,
		},
		"nil-some": {
			left:        nil,
			right:       DefaultInt64Set,
			wantUpdated: false,
			wantItems:   nil,
		},
		"some-nil": {
			left:        DefaultInt64Set,
			right:       nil,
			wantUpdated: false,
			wantItems:   DefaultInt64Set,
		},
		"empty-empty": {
			left:        EmptyInt64Set,
			right:       EmptyInt64Set,
			wantUpdated: false,
			wantItems:   EmptyInt64Set,
		},
		"empty-some": {
			left:        EmptyInt64Set,
			right:       DefaultInt64Set,
			wantUpdated: true,
			wantItems:   DefaultInt64Set,
		},
		"some-empty": {
			left:        DefaultInt64Set,
			right:       EmptyInt64Set,
			wantUpdated: false,
			wantItems:   DefaultInt64Set,
		},
		"same": {
			left:        DefaultInt64Set,
			right:       DefaultInt64Set,
			wantUpdated: false,
			wantItems:   DefaultInt64Set,
		},
		"intersection": {
			left:        set.New[int64](1, 2, 3),
			right:       set.New[int64](2, 4, 6),
			wantUpdated: true,
			wantItems:   set.New[int64](1, 2, 3, 4, 6),
		},
		"included": {
			left:        set.New[int64](1, 2, 3),
			right:       set.New[int64](2),
			wantUpdated: false,
			wantItems:   set.New[int64](1, 2, 3),
		},
		"no-intersection": {
			left:        set.New[int64](1, 2, 3),
			right:       set.New[int64](4, 6),
			wantUpdated: true,
			wantItems:   set.New[int64](1, 2, 3, 4, 6),
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.left)
		gotUpdated := set.Merge(set.Ptr(gotItems), testCase.right)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestMergeStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left        set.Set[Item]
		right       set.Set[Item]
		wantUpdated bool
		wantItems   set.Set[Item]
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:        nil,
			right:       nil,
			wantUpdated: false,
			wantItems:   nil,
		},
		"nil-empty": {
			left:        nil,
			right:       EmptyItemSet,
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty-nil": {
			left:        EmptyItemSet,
			right:       nil,
			wantUpdated: false,
			wantItems:   EmptyItemSet,
		},
		"nil-some": {
			left:        nil,
			right:       DefaultItemSet,
			wantUpdated: false,
			wantItems:   nil,
		},
		"some-nil": {
			left:        DefaultItemSet,
			right:       nil,
			wantUpdated: false,
			wantItems:   DefaultItemSet,
		},
		"empty-empty": {
			left:        EmptyItemSet,
			right:       EmptyItemSet,
			wantUpdated: false,
			wantItems:   EmptyItemSet,
		},
		"empty-some": {
			left:        EmptyItemSet,
			right:       DefaultItemSet,
			wantUpdated: true,
			wantItems:   DefaultItemSet,
		},
		"some-empty": {
			left:        DefaultItemSet,
			right:       EmptyItemSet,
			wantUpdated: false,
			wantItems:   DefaultItemSet,
		},
		"same": {
			left:        DefaultItemSet,
			right:       DefaultItemSet,
			wantUpdated: false,
			wantItems:   DefaultItemSet,
		},
		"intersection": {
			left:        set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}),
			right:       set.New(Item{Value: 2}, Item{Value: 4}, Item{Value: 6}),
			wantUpdated: true,
			wantItems:   set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}, Item{Value: 4}, Item{Value: 6}),
		},
		"included": {
			left:        set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}),
			right:       set.New(Item{Value: 2}),
			wantUpdated: false,
			wantItems:   set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}),
		},
		"no-intersection": {
			left:        set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}),
			right:       set.New(Item{Value: 4}, Item{Value: 6}),
			wantUpdated: true,
			wantItems:   set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}, Item{Value: 4}, Item{Value: 6}),
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.left)
		gotUpdated := set.Merge(set.Ptr(gotItems), testCase.right)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestMergeStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left        set.Set[*Item]
		right       set.Set[*Item]
		wantUpdated bool
		wantItems   set.Set[*Item]
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:        nil,
			right:       nil,
			wantUpdated: false,
			wantItems:   nil,
		},
		"nil-empty": {
			left:        nil,
			right:       EmptyItemPointerSet,
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty-nil": {
			left:        EmptyItemPointerSet,
			right:       nil,
			wantUpdated: false,
			wantItems:   EmptyItemPointerSet,
		},
		"nil-some": {
			left:        nil,
			right:       DefaultItemPointerSet,
			wantUpdated: false,
			wantItems:   nil,
		},
		"some-nil": {
			left:        DefaultItemPointerSet,
			right:       nil,
			wantUpdated: false,
			wantItems:   DefaultItemPointerSet,
		},
		"empty-empty": {
			left:        EmptyItemPointerSet,
			right:       EmptyItemPointerSet,
			wantUpdated: false,
			wantItems:   EmptyItemPointerSet,
		},
		"empty-some": {
			left:        EmptyItemPointerSet,
			right:       DefaultItemPointerSet,
			wantUpdated: true,
			wantItems:   DefaultItemPointerSet,
		},
		"some-empty": {
			left:        DefaultItemPointerSet,
			right:       EmptyItemPointerSet,
			wantUpdated: false,
			wantItems:   DefaultItemPointerSet,
		},
		"same": {
			left:        DefaultItemPointerSet,
			right:       DefaultItemPointerSet,
			wantUpdated: false,
			wantItems:   DefaultItemPointerSet,
		},
		"intersection": {
			left:        set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}),
			right:       set.New(&Item{Value: 2}, &Item{Value: 4}, &Item{Value: 6}),
			wantUpdated: true,
			wantItems:   set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}, &Item{Value: 4}, &Item{Value: 6}),
		},
		"included": {
			left:        set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}),
			right:       set.New(&Item{Value: 2}),
			wantUpdated: false,
			wantItems:   set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}),
		},
		"no-intersection": {
			left:        set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}),
			right:       set.New(&Item{Value: 4}, &Item{Value: 6}),
			wantUpdated: true,
			wantItems:   set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}, &Item{Value: 4}, &Item{Value: 6}),
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.left)
		gotUpdated := set.DeepMerge(set.Ptr(gotItems), testCase.right)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

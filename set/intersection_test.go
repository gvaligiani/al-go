package set_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
)

func TestIntersectionInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left      set.Set[int64]
		right     set.Set[int64]
		wantItems set.Set[int64]
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:      nil,
			right:     nil,
			wantItems: nil,
		},
		"nil-empty": {
			left:      nil,
			right:     EmptyInt64Set,
			wantItems: EmptyInt64Set,
		},
		"empty-nil": {
			left:      EmptyInt64Set,
			right:     nil,
			wantItems: EmptyInt64Set,
		},
		"nil-some": {
			left:      nil,
			right:     DefaultInt64Set,
			wantItems: EmptyInt64Set,
		},
		"some-nil": {
			left:      DefaultInt64Set,
			right:     nil,
			wantItems: EmptyInt64Set,
		},
		"empty-empty": {
			left:      EmptyInt64Set,
			right:     EmptyInt64Set,
			wantItems: EmptyInt64Set,
		},
		"empty-some": {
			left:      EmptyInt64Set,
			right:     DefaultInt64Set,
			wantItems: EmptyInt64Set,
		},
		"some-empty": {
			left:      DefaultInt64Set,
			right:     EmptyInt64Set,
			wantItems: EmptyInt64Set,
		},
		"same": {
			left:      DefaultInt64Set,
			right:     DefaultInt64Set,
			wantItems: DefaultInt64Set,
		},
		"intersection": {
			left:      set.New[int64](1, 2, 3),
			right:     set.New[int64](2, 4, 6),
			wantItems: set.New[int64](2),
		},
		"included": {
			left:      set.New[int64](1, 2, 3),
			right:     set.New[int64](2),
			wantItems: set.New[int64](2),
		},
		"no-intersection": {
			left:      set.New[int64](1, 2, 3),
			right:     set.New[int64](4, 6),
			wantItems: EmptyInt64Set,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Intersection(testCase.left, testCase.right)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestIntersectionStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left      set.Set[Item]
		right     set.Set[Item]
		wantItems set.Set[Item]
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:      nil,
			right:     nil,
			wantItems: nil,
		},
		"nil-empty": {
			left:      nil,
			right:     EmptyItemSet,
			wantItems: EmptyItemSet,
		},
		"empty-nil": {
			left:      EmptyItemSet,
			right:     nil,
			wantItems: EmptyItemSet,
		},
		"nil-some": {
			left:      nil,
			right:     DefaultItemSet,
			wantItems: EmptyItemSet,
		},
		"some-nil": {
			left:      DefaultItemSet,
			right:     nil,
			wantItems: EmptyItemSet,
		},
		"empty-empty": {
			left:      EmptyItemSet,
			right:     EmptyItemSet,
			wantItems: EmptyItemSet,
		},
		"empty-some": {
			left:      EmptyItemSet,
			right:     DefaultItemSet,
			wantItems: EmptyItemSet,
		},
		"some-empty": {
			left:      DefaultItemSet,
			right:     EmptyItemSet,
			wantItems: EmptyItemSet,
		},
		"same": {
			left:      DefaultItemSet,
			right:     DefaultItemSet,
			wantItems: DefaultItemSet,
		},
		"intersection": {
			left:      set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}),
			right:     set.New(Item{Value: 2}, Item{Value: 4}, Item{Value: 6}),
			wantItems: set.New(Item{Value: 2}),
		},
		"included": {
			left:      set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}),
			right:     set.New(Item{Value: 2}),
			wantItems: set.New(Item{Value: 2}),
		},
		"no-intersection": {
			left:      set.New(Item{Value: 1}, Item{Value: 2}, Item{Value: 3}),
			right:     set.New(Item{Value: 4}, Item{Value: 6}),
			wantItems: EmptyItemSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Intersection(testCase.left, testCase.right)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestIntersectionStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left      set.Set[*Item]
		right     set.Set[*Item]
		wantItems set.Set[*Item]
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:      nil,
			right:     nil,
			wantItems: nil,
		},
		"nil-empty": {
			left:      nil,
			right:     EmptyItemPointerSet,
			wantItems: EmptyItemPointerSet,
		},
		"empty-nil": {
			left:      EmptyItemPointerSet,
			right:     nil,
			wantItems: EmptyItemPointerSet,
		},
		"nil-some": {
			left:      nil,
			right:     DefaultItemPointerSet,
			wantItems: EmptyItemPointerSet,
		},
		"some-nil": {
			left:      DefaultItemPointerSet,
			right:     nil,
			wantItems: EmptyItemPointerSet,
		},
		"empty-empty": {
			left:      EmptyItemPointerSet,
			right:     EmptyItemPointerSet,
			wantItems: EmptyItemPointerSet,
		},
		"empty-some": {
			left:      EmptyItemPointerSet,
			right:     DefaultItemPointerSet,
			wantItems: EmptyItemPointerSet,
		},
		"some-empty": {
			left:      DefaultItemPointerSet,
			right:     EmptyItemPointerSet,
			wantItems: EmptyItemPointerSet,
		},
		"same": {
			left:      DefaultItemPointerSet,
			right:     DefaultItemPointerSet,
			wantItems: DefaultItemPointerSet,
		},
		"intersection": {
			left:      set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}),
			right:     set.New(&Item{Value: 2}, &Item{Value: 4}, &Item{Value: 6}),
			wantItems: set.New(&Item{Value: 2}),
		},
		"included": {
			left:      set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}),
			right:     set.New(&Item{Value: 2}),
			wantItems: set.New(&Item{Value: 2}),
		},
		"no-intersection": {
			left:      set.New(&Item{Value: 1}, &Item{Value: 2}, &Item{Value: 3}),
			right:     set.New(&Item{Value: 4}, &Item{Value: 6}),
			wantItems: EmptyItemPointerSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.DeepIntersection(testCase.left, testCase.right)

		// assert
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
)

func TestEqualsInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       set.Set[int64]
		right      set.Set[int64]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyInt64Set,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyInt64Set,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyInt64Set,
			right:      EmptyInt64Set,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultInt64Set,
			right:      EmptyInt64Set,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyInt64Set,
			right:      DefaultInt64Set,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultInt64Set,
			right:      DefaultInt64Set,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultInt64Set,
			right:      OtherInt64Set,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherInt64Set,
			right:      DefaultInt64Set,
			wantEquals: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		left := testCase.left.Copy()
		right := testCase.right.Copy()
		gotEquals := set.Equal(left, right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

func TestEqualsStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       set.Set[Item]
		right      set.Set[Item]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyItemSet,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyItemSet,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyItemSet,
			right:      EmptyItemSet,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultItemSet,
			right:      EmptyItemSet,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyItemSet,
			right:      DefaultItemSet,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultItemSet,
			right:      DefaultItemSet,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultItemSet,
			right:      OtherItemSet,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherItemSet,
			right:      DefaultItemSet,
			wantEquals: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		left := testCase.left.Copy()
		right := testCase.right.Copy()
		gotEquals := set.Equal(left, right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

func TestEqualsStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       set.Set[*Item]
		right      set.Set[*Item]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyItemPointerSet,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyItemPointerSet,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyItemPointerSet,
			right:      EmptyItemPointerSet,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultItemPointerSet,
			right:      EmptyItemPointerSet,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyItemPointerSet,
			right:      DefaultItemPointerSet,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultItemPointerSet,
			right:      DefaultItemPointerSet,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultItemPointerSet,
			right:      OtherItemPointerSet,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherItemPointerSet,
			right:      DefaultItemPointerSet,
			wantEquals: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		left := testCase.left.Copy()
		right := testCase.right.Copy()
		gotEquals := set.Equal(left, right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestEqualsInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       dict.Dict[int, int64]
		right      dict.Dict[int, int64]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyInt64Dict,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyInt64Dict,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyInt64Dict,
			right:      EmptyInt64Dict,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultInt64Dict,
			right:      EmptyInt64Dict,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyInt64Dict,
			right:      DefaultInt64Dict,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultInt64Dict,
			right:      DefaultInt64Dict,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultInt64Dict,
			right:      OtherInt64Dict,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherInt64Dict,
			right:      DefaultInt64Dict,
			wantEquals: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotEquals := dict.Equal(testCase.left, testCase.right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

func TestEqualsStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       dict.Dict[int, Item]
		right      dict.Dict[int, Item]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyItemDict,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyItemDict,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyItemDict,
			right:      EmptyItemDict,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultItemDict,
			right:      EmptyItemDict,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyItemDict,
			right:      DefaultItemDict,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultItemDict,
			right:      DefaultItemDict,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultItemDict,
			right:      OtherItemDict,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherItemDict,
			right:      DefaultItemDict,
			wantEquals: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotEquals := dict.Equal(testCase.left, testCase.right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

func TestEqualsStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       dict.Dict[int, *Item]
		right      dict.Dict[int, *Item]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyItemPointerDict,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyItemPointerDict,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyItemPointerDict,
			right:      EmptyItemPointerDict,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultItemPointerDict,
			right:      EmptyItemPointerDict,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyItemPointerDict,
			right:      DefaultItemPointerDict,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultItemPointerDict,
			right:      DefaultItemPointerDict,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultItemPointerDict,
			right:      OtherItemPointerDict,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherItemPointerDict,
			right:      DefaultItemPointerDict,
			wantEquals: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotEquals := dict.Equal(testCase.left, testCase.right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

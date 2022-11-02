package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestEqualsInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       list.List[int64]
		right      list.List[int64]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyInt64List,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyInt64List,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyInt64List,
			right:      EmptyInt64List,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultInt64List,
			right:      EmptyInt64List,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyInt64List,
			right:      DefaultInt64List,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultInt64List,
			right:      DefaultInt64List,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultInt64List,
			right:      OtherInt64List,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherInt64List,
			right:      DefaultInt64List,
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
		gotEquals := list.Equals(left, right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

func TestEqualsStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       list.List[Item]
		right      list.List[Item]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyItemList,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyItemList,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyItemList,
			right:      EmptyItemList,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultItemList,
			right:      EmptyItemList,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyItemList,
			right:      DefaultItemList,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultItemList,
			right:      DefaultItemList,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultItemList,
			right:      OtherItemList,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherItemList,
			right:      DefaultItemList,
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
		gotEquals := list.Equals(left, right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

func TestEqualsStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		left       list.List[*Item]
		right      list.List[*Item]
		wantEquals bool
	}

	testCases := map[string]TestCase{
		"nil-nil": {
			left:       nil,
			right:      nil,
			wantEquals: true,
		},
		"empty-nil": {
			left:       EmptyItemPointerList,
			right:      nil,
			wantEquals: false,
		},
		"nil-empty": {
			left:       nil,
			right:      EmptyItemPointerList,
			wantEquals: false,
		},
		"empty-empty": {
			left:       EmptyItemPointerList,
			right:      EmptyItemPointerList,
			wantEquals: true,
		},
		"default-empty": {
			left:       DefaultItemPointerList,
			right:      EmptyItemPointerList,
			wantEquals: false,
		},
		"empty-default": {
			left:       EmptyItemPointerList,
			right:      DefaultItemPointerList,
			wantEquals: false,
		},
		"default-default": {
			left:       DefaultItemPointerList,
			right:      DefaultItemPointerList,
			wantEquals: true,
		},
		"default-other": {
			left:       DefaultItemPointerList,
			right:      OtherItemPointerList,
			wantEquals: false,
		},
		"other-default": {
			left:       OtherItemPointerList,
			right:      DefaultItemPointerList,
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
		gotEquals := list.Equals(left, right)

		// assert
		require.Equalf(t, testCase.wantEquals, gotEquals, "wrong equals!")
	})
}

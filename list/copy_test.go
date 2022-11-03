package list_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestCopyInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[int64]
		wantItems list.List[int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyInt64List,
			wantItems: EmptyInt64List,
		},
		"default": {
			items:     DefaultInt64List,
			wantItems: DefaultInt64List,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[Item]
		wantItems list.List[Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemList,
			wantItems: EmptyItemList,
		},
		"default": {
			items:     DefaultItemList,
			wantItems: DefaultItemList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[*Item]
		wantItems list.List[*Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemPointerList,
			wantItems: EmptyItemPointerList,
		},
		"default": {
			items:     DefaultItemPointerList,
			wantItems: DefaultItemPointerList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)

		// assert
		assertEquals(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

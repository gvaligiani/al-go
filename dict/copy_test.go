package dict_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestCopyInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, int64]
		wantItems dict.Dict[int, int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyInt64Dict,
			wantItems: EmptyInt64Dict,
		},
		"default": {
			items:     DefaultInt64Dict,
			wantItems: DefaultInt64Dict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.Copy(testCase.items)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, Item]
		wantItems dict.Dict[int, Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemDict,
			wantItems: EmptyItemDict,
		},
		"default": {
			items:     DefaultItemDict,
			wantItems: DefaultItemDict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.Copy(testCase.items)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, *Item]
		wantItems dict.Dict[int, *Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemPointerDict,
			wantItems: EmptyItemPointerDict,
		},
		"default": {
			items:     DefaultItemPointerDict,
			wantItems: SameItemPointerDict,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := dict.Copy(testCase.items)

		// assert
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

package set_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
)

func TestCopyInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[int64]
		wantItems set.Set[int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyInt64Set,
			wantItems: EmptyInt64Set,
		},
		"default": {
			items:     DefaultInt64Set,
			wantItems: DefaultInt64Set,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[Item]
		wantItems set.Set[Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemSet,
			wantItems: EmptyItemSet,
		},
		"default": {
			items:     DefaultItemSet,
			wantItems: DefaultItemSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[*Item]
		wantItems set.Set[*Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemPointerSet,
			wantItems: EmptyItemPointerSet,
		},
		"default": {
			items:     DefaultItemPointerSet,
			wantItems: SameItemPointerSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.Copy(testCase.items)

		// assert
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

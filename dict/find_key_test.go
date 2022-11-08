package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/test"
)

func TestFindByKeyInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, int64]
		key       int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			key:       30,
			wantFound: false,
		},
		"empty": {
			items:     EmptyInt64Dict,
			key:       30,
			wantFound: false,
		},
		"found": {
			items:     DefaultInt64Dict,
			key:       30,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultInt64Dict,
			key:       90,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindKey(testCase.items, testCase.key)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindByKeyStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, Item]
		key       int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			key:       30,
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemDict,
			key:       30,
			wantFound: false,
		},
		"found": {
			items:     DefaultItemDict,
			key:       30,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemDict,
			key:       90,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindKey(testCase.items, testCase.key)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindByKeyStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     dict.Dict[int, *Item]
		key       int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			key:       30,
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemPointerDict,
			key:       30,
			wantFound: false,
		},
		"found": {
			items:     DefaultItemPointerDict,
			key:       30,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemPointerDict,
			key:       90,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindKey(testCase.items, testCase.key)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

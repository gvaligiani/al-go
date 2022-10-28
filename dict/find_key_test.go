package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestFindKeyInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		key     int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			key:     30,
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyInt64Dict,
			key:     30,
			wantFound: false,
		},
		"found": {
			items:     test.Int64Dict,
			key:     30,
			wantFound: true,
		},
		"not-found": {
			items:     test.Int64Dict,
			key:     90,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindKey[int,int64](testCase.items, testCase.key)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindKeyStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]test.Item
		key     int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			key:     30,
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemDict,
			key:     30,
			wantFound: false,
		},
		"found": {
			items:     test.ItemDict,
			key:     30,
			wantFound: true,
		},
		"not-found": {
			items:     test.ItemDict,
			key:     90,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindKey[int,test.Item](testCase.items, testCase.key)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindKeyStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]*test.Item
		key     int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			key:     30,
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemPointerDict,
			key:     30,
			wantFound: false,
		},
		"found": {
			items:     test.ItemPointerDict,
			key:     30,
			wantFound: true,
		},
		"not-found": {
			items:     test.ItemPointerDict,
			key:     90,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindKey[int, *test.Item](testCase.items, testCase.key)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

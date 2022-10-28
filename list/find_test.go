package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/list"
	"github.com/gvaligiani/algo/test"
)

func TestFindInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		value     int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     34,
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyInt64List,
			value:     34,
			wantFound: false,
		},
		"found": {
			items:     test.DefaultInt64List,
			value:     34,
			wantFound: true,
		},
		"not-found": {
			items:     test.DefaultInt64List,
			value:     34,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := list.Find[int64](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []test.Item
		value     test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     test.Item{Value: 34},
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemList,
			value:     test.Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     test.DefaultItemList,
			value:     test.Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     test.DefaultItemList,
			value:     test.Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := list.Find[test.Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []*test.Item
		value     *test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     &test.Item{Value: 34},
			wantFound: false,
		},
		"empty": {
			items:     test.EmptyItemPointerList,
			value:     &test.Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     test.DefaultItemPointerList,
			value:     &test.Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     test.DefaultItemPointerList,
			value:     &test.Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := list.Find[*test.Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

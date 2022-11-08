package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
)

func TestFindIndexInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []int64
		value     int64
		wantIndex int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     34,
			wantIndex: -1,
			wantFound: false,
		},
		"empty": {
			items:     EmptyInt64List,
			value:     34,
			wantIndex: -1,
			wantFound: false,
		},
		"found": {
			items:     DefaultInt64List,
			value:     34,
			wantIndex: 2,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultInt64List,
			value:     99,
			wantIndex: -1,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotFound := list.FindIndexFromValue(testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindIndexStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[Item]
		value     Item
		wantIndex int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     Item{Value: 34},
			wantIndex: -1,
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemList,
			value:     Item{Value: 34},
			wantIndex: -1,
			wantFound: false,
		},
		"found": {
			items:     DefaultItemList,
			value:     Item{Value: 34},
			wantIndex: 2,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemList,
			value:     Item{Value: 99},
			wantIndex: -1,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotFound := list.FindIndexFromValue(testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindIndexStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     list.List[*Item]
		value     *Item
		wantIndex int
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     &Item{Value: 34},
			wantIndex: -1,
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemPointerList,
			value:     &Item{Value: 34},
			wantIndex: -1,
			wantFound: false,
		},
		"found": {
			items:     DefaultItemPointerList,
			value:     &Item{Value: 34},
			wantIndex: 2,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemPointerList,
			value:     &Item{Value: 99},
			wantIndex: -1,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotIndex, gotFound := list.DeepFindIndexFromValue(testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantIndex, gotIndex, "wrong index!")
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

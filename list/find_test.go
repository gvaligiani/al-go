package list_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
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
			items:     EmptyInt64List,
			value:     34,
			wantFound: false,
		},
		"found": {
			items:     DefaultInt64List,
			value:     34,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultInt64List,
			value:     34,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

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
		items     []Item
		value     Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     Item{Value: 34},
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemList,
			value:     Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     DefaultItemList,
			value:     Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemList,
			value:     Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := list.Find[Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     []*Item
		value     *Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     &Item{Value: 34},
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemPointerList,
			value:     &Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     DefaultItemPointerList,
			value:     &Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemPointerList,
			value:     &Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := list.Find[*Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

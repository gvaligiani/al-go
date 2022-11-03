package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
)

func TestFindInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[int64]
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
			items:     EmptyInt64Set,
			value:     34,
			wantFound: false,
		},
		"found": {
			items:     DefaultInt64Set,
			value:     34,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultInt64Set,
			value:     34,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find(testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[Item]
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
			items:     EmptyItemSet,
			value:     Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     DefaultItemSet,
			value:     Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemSet,
			value:     Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find(testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[*Item]
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
			items:     EmptyItemPointerSet,
			value:     &Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     DefaultItemPointerSet,
			value:     &Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemPointerSet,
			value:     &Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find(testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

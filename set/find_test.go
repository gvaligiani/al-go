package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestFindInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int64]struct{}
		value     int64
		wantFound bool
	}

	empty := map[int64]struct{}{}
	items := map[int64]struct{}{
		21: {},
		12: {},
		34: {},
		87: {},
		52: {},
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     34,
			wantFound: false,
		},
		"empty": {
			items:     empty,
			value:     34,
			wantFound: false,
		},
		"found": {
			items:     items,
			value:     34,
			wantFound: true,
		},
		"not-found": {
			items:     items,
			value:     34,
			wantFound: true,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find[int64](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items     map[Item]struct{}
		value     Item
		wantFound bool
	}

	empty := map[Item]struct{}{}
	items := map[Item]struct{}{
		{value: 21}: {},
		{value: 12}: {},
		{value: 34}: {},
		{value: 87}: {},
		{value: 52}: {},
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     Item{value: 34},
			wantFound: false,
		},
		"empty": {
			items:     empty,
			value:     Item{value: 34},
			wantFound: false,
		},
		"found": {
			items:     items,
			value:     Item{value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     items,
			value:     Item{value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find[Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items     map[*Item]struct{}
		value     *Item
		wantFound bool
	}

	empty := map[*Item]struct{}{}
	items := map[*Item]struct{}{
		{value: 21}: {},
		{value: 12}: {},
		{value: 34}: {},
		{value: 87}: {},
		{value: 52}: {},
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     &Item{value: 34},
			wantFound: false,
		},
		"empty": {
			items:     empty,
			value:     &Item{value: 34},
			wantFound: false,
		},
		"found": {
			items:     items,
			value:     &Item{value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     items,
			value:     &Item{value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := set.Find[*Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

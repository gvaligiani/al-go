package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestFindIfNotInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int64]struct{}
		predicate func(int64) bool
		wantItem  int64
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
			predicate: func(i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"empty": {
			items:     empty,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"no-match": {
			items:     items,
			predicate: func(i int64) bool { return i < 100 },
			wantFound: false,
			wantItem:  0,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(i int64) bool { return i%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  21,
		// },
		"one-match": {
			items:     items,
			predicate: func(i int64) bool { return i < 60 },
			wantFound: true,
			wantItem:  87,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfNotStruct(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items     map[Item]struct{}
		predicate func(Item) bool
		wantItem  Item
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
			predicate: func(i Item) bool { return i.value%10 == 3 },
			wantFound: false,
			wantItem:  Item{},
		},
		"empty": {
			items:     empty,
			predicate: func(i Item) bool { return i.value%10 == 3 },
			wantFound: false,
			wantItem:  Item{},
		},
		"no-match": {
			items:     items,
			predicate: func(i Item) bool { return i.value < 100 },
			wantFound: false,
			wantItem:  Item{},
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(i Item) bool { return i.value%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  Item{value: 21},
		// },
		"one-match": {
			items:     items,
			predicate: func(i Item) bool { return i.value < 60 },
			wantFound: true,
			wantItem:  Item{value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot[Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfNotStructPointer(t *testing.T) {

	//
	// test cases
	//

	type Item struct {
		value int64
	}

	type TestCase struct {
		items     map[*Item]struct{}
		predicate func(*Item) bool
		wantItem  *Item
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
			predicate: func(i *Item) bool { return i.value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"empty": {
			items:     empty,
			predicate: func(i *Item) bool { return i.value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"no-match": {
			items:     items,
			predicate: func(i *Item) bool { return i.value < 100 },
			wantFound: false,
			wantItem:  nil,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     items,
		// 	predicate: func(i *Item) bool { return i.value%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  &Item{value: 21},
		// },
		"one-match": {
			items:     items,
			predicate: func(i *Item) bool { return i.value < 60 },
			wantFound: true,
			wantItem:  &Item{value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot[*Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

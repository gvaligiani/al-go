package list_test

import (
	"math/rand"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/list"
	"github.com/gvaligiani/al.go/test"
	"github.com/stretchr/testify/require"
)

func TestCheckIndexes(t *testing.T) {
	rand.Seed(time.Now().UnixMicro())
	type Obj struct {
		index int
		value int
	}
	objs := list.List[Obj]{}
	for i := 0; i < rand.Intn(100); i++ {
		objs.Add(Obj{index: i, value: rand.Intn(10)})
	}
	objs.RemoveIf(func(i int, o Obj) bool {
		require.Equal(t, i, o.index, "wrong index")
		return o.value%2 == 1
	})
}

func TestRemoveIfInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       list.List[int64]
		predicate   list.Predicate[int64]
		wantUpdated bool
		wantItems   list.List[int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(_ int, i int64) bool { return i%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyInt64List,
			predicate:   func(_ int, i int64) bool { return i%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyInt64List,
		},
		"remove-none": {
			items:       DefaultInt64List,
			predicate:   func(_ int, i int64) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultInt64List,
		},
		"remove-odd": {
			items:       DefaultInt64List,
			predicate:   func(_ int, i int64) bool { return i%2 == 0 },
			wantUpdated: true,
			wantItems: list.New[int64](
				21,
				87,
			),
		},
		"remove-even": {
			items:       DefaultInt64List,
			predicate:   func(_ int, i int64) bool { return i%2 == 1 },
			wantUpdated: true,
			wantItems: list.New[int64]( // list re-shuffled
				52,
				12,
				34,
			),
		},
		"remove-all": {
			items:       DefaultInt64List,
			predicate:   func(_ int, i int64) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyInt64List,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		gotUpdated := list.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestRemoveIfStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       list.List[Item]
		predicate   list.Predicate[Item]
		wantUpdated bool
		wantItems   list.List[Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyItemList,
			predicate:   func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyItemList,
		},
		"remove-none": {
			items:       DefaultItemList,
			predicate:   func(_ int, item Item) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultItemList,
		},
		"remove-odd": {
			items:       DefaultItemList,
			predicate:   func(_ int, item Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: list.List[Item]{
				Item{Value: 21},
				Item{Value: 87},
			},
		},
		"remove-even": {
			items:       DefaultItemList,
			predicate:   func(_ int, item Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: list.List[Item]{ // list re-shuffled
				Item{Value: 52},
				Item{Value: 12},
				Item{Value: 34},
			},
		},
		"remove-all": {
			items:       DefaultItemList,
			predicate:   func(_ int, item Item) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyItemList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		gotUpdated := list.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestRemoveIfStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items       list.List[*Item]
		predicate   list.Predicate[*Item]
		wantUpdated bool
		wantItems   list.List[*Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:       nil,
			predicate:   func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   nil,
		},
		"empty": {
			items:       EmptyItemPointerList,
			predicate:   func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: false,
			wantItems:   EmptyItemPointerList,
		},
		"remove-none": {
			items:       DefaultItemPointerList,
			predicate:   func(_ int, item *Item) bool { return false },
			wantUpdated: false,
			wantItems:   DefaultItemPointerList,
		},
		"remove-odd": {
			items:       DefaultItemPointerList,
			predicate:   func(_ int, item *Item) bool { return item.Value%2 == 0 },
			wantUpdated: true,
			wantItems: list.List[*Item]{
				&Item{Value: 21},
				&Item{Value: 87},
			},
		},
		"remove-even": {
			items:       DefaultItemPointerList,
			predicate:   func(_ int, item *Item) bool { return item.Value%2 == 1 },
			wantUpdated: true,
			wantItems: list.List[*Item]{ // list reshuffled
				&Item{Value: 52},
				&Item{Value: 12},
				&Item{Value: 34},
			},
		},
		"remove-all": {
			items:       DefaultItemPointerList,
			predicate:   func(_ int, item *Item) bool { return true },
			wantUpdated: true,
			wantItems:   EmptyItemPointerList,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := list.Copy(testCase.items)
		gotUpdated := list.RemoveIf(&gotItems, testCase.predicate)

		// assert
		require.Equal(t, testCase.wantUpdated, gotUpdated, "wrong updated!")
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

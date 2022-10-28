package test

// alias
type Int64List []int64
type Int64Set map[int64]struct{}
type ItemList []Item
type ItemSet map[Item]struct{}

// int64
var (
	EmptyInt64Set   = map[int64]struct{}{}
	DefaultInt64Set = map[int64]struct{}{
		21: {},
		12: {},
		34: {},
		87: {},
		52: {},
	}

	EmptyInt64Dict   = map[int]int64{}
	DefaultInt64Dict = map[int]int64{
		10: 21,
		20: 12,
		30: 34,
		40: 87,
		50: 52,
	}

	EmptyInt64List   = []int64{}
	DefaultInt64List = []int64{
		21,
		12,
		34,
		87,
		52,
	}
)

// Item

type Item struct {
	Value int64
}

var (
	EmptyItemSet   = map[Item]struct{}{}
	DefaultItemSet = map[Item]struct{}{
		Item{Value: 21}: {},
		Item{Value: 12}: {},
		Item{Value: 34}: {},
		Item{Value: 87}: {},
		Item{Value: 52}: {},
	}

	EmptyItemDict   = map[int]Item{}
	DefaultItemDict = map[int]Item{
		10: Item{Value: 21},
		20: Item{Value: 12},
		30: Item{Value: 34},
		40: Item{Value: 87},
		50: Item{Value: 52},
	}

	EmptyItemList   = []Item{}
	DefaultItemList = []Item{
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 34},
		Item{Value: 87},
		Item{Value: 52},
	}
)

// *test.Item

var (
	EmptyItemPointerSet   = map[*Item]struct{}{}
	DefaultItemPointerSet = map[*Item]struct{}{
		&Item{Value: 21}: {},
		&Item{Value: 12}: {},
		&Item{Value: 34}: {},
		&Item{Value: 87}: {},
		&Item{Value: 52}: {},
	}

	EmptyItemPointerDict   = map[int]*Item{}
	DefaultItemPointerDict = map[int]*Item{
		10: &Item{Value: 21},
		20: &Item{Value: 12},
		30: &Item{Value: 34},
		40: &Item{Value: 87},
		50: &Item{Value: 52},
	}

	EmptyItemPointerList   = []*Item{}
	DefaultItemPointerList = []*Item{
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 34},
		&Item{Value: 87},
		&Item{Value: 52},
	}
)

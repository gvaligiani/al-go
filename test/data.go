package test

// int64
var (
	EmptyInt64Set = map[int64]struct{}{}
	Int64Set      = map[int64]struct{}{
		21: {},
		12: {},
		34: {},
		87: {},
		52: {},
	}

	EmptyInt64Dict = map[int]int64{}
	Int64Dict      = map[int]int64{
		10: 21,
		20: 12,
		30: 34,
		40: 87,
		50: 52,
	}

	EmptyInt64List = []int64{}
	Int64List      = []int64{
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
	EmptyItemSet = map[Item]struct{}{}
	ItemSet      = map[Item]struct{}{
		Item{Value: 21}: {},
		Item{Value: 12}: {},
		Item{Value: 34}: {},
		Item{Value: 87}: {},
		Item{Value: 52}: {},
	}

	EmptyItemDict = map[int]Item{}
	ItemDict      = map[int]Item{
		10: Item{Value: 21},
		20: Item{Value: 12},
		30: Item{Value: 34},
		40: Item{Value: 87},
		50: Item{Value: 52},
	}

	EmptyItemList = []Item{}
	ItemList      = []Item{
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 34},
		Item{Value: 87},
		Item{Value: 52},
	}
)

// *test.Item

var (
	EmptyItemPointerSet = map[*Item]struct{}{}
	ItemPointerSet      = map[*Item]struct{}{
		&Item{Value: 21}: {},
		&Item{Value: 12}: {},
		&Item{Value: 34}: {},
		&Item{Value: 87}: {},
		&Item{Value: 52}: {},
	}

	EmptyItemPointerDict = map[int]*Item{}
	ItemPointerDict      = map[int]*Item{
		10: &Item{Value: 21},
		20: &Item{Value: 12},
		30: &Item{Value: 34},
		40: &Item{Value: 87},
		50: &Item{Value: 52},
	}

	EmptyItemPointerList = []*Item{}
	ItemPointerList      = []*Item{
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 34},
		&Item{Value: 87},
		&Item{Value: 52},
	}
)

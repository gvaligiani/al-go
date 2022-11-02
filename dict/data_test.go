package dict_test

// int64
var (
	EmptyInt64Dict   = map[int]int64{}
	DefaultInt64Dict = map[int]int64{
		10: 21,
		20: 12,
		30: 34,
		40: 87,
		50: 52,
	}
)

// Item

type Item struct {
	Value int64
}

var (
	EmptyItemDict   = map[int]Item{}
	DefaultItemDict = map[int]Item{
		10: Item{Value: 21},
		20: Item{Value: 12},
		30: Item{Value: 34},
		40: Item{Value: 87},
		50: Item{Value: 52},
	}
)

// *Item

var (
	EmptyItemPointerDict   = map[int]*Item{}
	DefaultItemPointerDict = map[int]*Item{
		10: &Item{Value: 21},
		20: &Item{Value: 12},
		30: &Item{Value: 34},
		40: &Item{Value: 87},
		50: &Item{Value: 52},
	}
)

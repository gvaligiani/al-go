package list_test

// int64
var (
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
	EmptyItemList   = []Item{}
	DefaultItemList = []Item{
		Item{Value: 21},
		Item{Value: 12},
		Item{Value: 34},
		Item{Value: 87},
		Item{Value: 52},
	}
)

// *Item

var (
	EmptyItemPointerList   = []*Item{}
	DefaultItemPointerList = []*Item{
		&Item{Value: 21},
		&Item{Value: 12},
		&Item{Value: 34},
		&Item{Value: 87},
		&Item{Value: 52},
	}
)

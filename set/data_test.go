package set_test

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
)

// *Item

var (
	EmptyItemPointerSet   = map[*Item]struct{}{}
	DefaultItemPointerSet = map[*Item]struct{}{
		&Item{Value: 21}: {},
		&Item{Value: 12}: {},
		&Item{Value: 34}: {},
		&Item{Value: 87}: {},
		&Item{Value: 52}: {},
	}
)

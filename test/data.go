package test

// int64

type Int64SetAlias map[int64]struct{}

var (
	EmptyInt64Set = Int64SetAlias{}
	Int64Set      = Int64SetAlias{
		21: {},
		12: {},
		34: {},
		87: {},
		52: {},
	}
)

// item

type Item struct {
	value int64
}

type ItemSetAlias map[Item]struct{}

var (
	EmptyItemSet = ItemSetAlias{}
	ItemSet      = ItemSetAlias{
		Item{value: 21}: {},
		Item{value: 12}: {},
		Item{value: 34}: {},
		Item{value: 87}: {},
		Item{value: 52}: {},
	}
)

// *Item

type ItemPointerSetAlias map[*Item]struct{}

var (
	EmptyItemPointerSet = ItemPointerSetAlias{}
	ItemPointerSet      = ItemPointerSetAlias{
		&Item{value: 21}: {},
		&Item{value: 12}: {},
		&Item{value: 34}: {},
		&Item{value: 87}: {},
		&Item{value: 52}: {},
	}
)

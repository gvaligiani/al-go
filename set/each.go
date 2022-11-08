package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func Each[T comparable, S ~map[T]struct{}](items S, consumer util.Consumer[T]) {
	dict.EachKey(items, util.ConsumeOnFirstArg[T, struct{}](consumer))
}

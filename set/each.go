package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func Each[V comparable, S ~map[V]struct{}](s S, consumer util.Consumer[V]) {
	dict.EachKey(s, util.ConsumeOnFirstArg[V, struct{}](consumer))
}

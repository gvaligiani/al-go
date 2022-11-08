package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func NoneOf[V comparable, S ~map[V]struct{}](s S, predicate util.Predicate[V]) bool {
	return dict.NoKeyOf(s, util.TestOnFirstArg[V, struct{}](predicate))
}

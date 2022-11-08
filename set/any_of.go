package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func AnyOf[V comparable, S ~map[V]struct{}](s S, predicate util.Predicate[V]) bool {
	return dict.AnyKeyOf(s, util.TestOnFirstArg[V, struct{}](predicate))
}

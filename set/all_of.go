package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func AllOf[T comparable, S ~map[T]struct{}](items S, predicate util.Predicate[T]) bool {
	return dict.AllKeyOf(items, util.TestOnFirstArg[T, struct{}](predicate))
}

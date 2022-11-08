package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func FindIf[T comparable, S ~map[T]struct{}](items S, predicate util.Predicate[T]) (T, bool) {
	item, _, found := dict.FindKeyIf(items, util.TestOnFirstArg[T, struct{}](predicate))
	return item, found
}

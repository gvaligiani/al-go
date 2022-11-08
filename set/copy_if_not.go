package set

import "github.com/gvaligiani/al.go/util"

func CopyIfNot[T comparable, S ~map[T]struct{}](items S, predicate util.Predicate[T]) S {
	return CopyIf(items, util.Not(predicate))
}

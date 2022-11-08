package set

import "github.com/gvaligiani/al.go/util"

func FindIfNot[T comparable, S ~map[T]struct{}](items S, predicate util.Predicate[T]) (T, bool) {
	return FindIf(items, util.Not(predicate))
}

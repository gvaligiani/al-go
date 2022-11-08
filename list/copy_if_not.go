package list

import "github.com/gvaligiani/al.go/util"

func CopyIfNot[T any, L ~[]T](items L, predicate util.Predicate[T]) L {
	return CopyIf(items, util.Not(predicate))
}

func CopyIndexIfNot[T any, L ~[]T](items L, predicate util.BiPredicate[int, T]) L {
	return CopyIndexIf(items, util.BiNot(predicate))
}

package list

import "github.com/gvaligiani/al.go/util"

func KeepIf[T any, L ~[]T](items *L, predicate util.Predicate[T]) bool {
	return RemoveIf(items, util.Not(predicate))
}

func KeepIndexIf[T any, L ~[]T](items *L, predicate util.BiPredicate[int, T]) bool {
	return RemoveIndexIf(items, util.BiNot(predicate))
}

package dict

import "github.com/gvaligiani/al.go/util"

func KeepIf[K comparable, T any, M ~map[K]T](items *M, predicate util.Predicate[T]) bool {
	return RemoveIf(items, util.Not(predicate))
}

func KeepKeyIf[K comparable, T any, M ~map[K]T](items *M, predicate util.BiPredicate[K, T]) bool {
	return RemoveKeyIf(items, util.BiNot(predicate))
}

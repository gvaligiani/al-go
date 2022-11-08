package list

import "github.com/gvaligiani/al.go/util"

func FindIfNot[T any, L ~[]T](items L, predicate util.Predicate[T]) (T, bool) {
	return FindIf(items, util.Not(predicate))
}

func FindIndexIfNot[T any, L ~[]T](items L, predicate util.BiPredicate[int, T]) (int, T, bool) {
	return FindIndexIf(items, util.BiNot(predicate))
}

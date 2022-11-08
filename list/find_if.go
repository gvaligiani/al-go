package list

import "github.com/gvaligiani/al.go/util"

func FindIf[T any, L ~[]T](items L, predicate util.Predicate[T]) (T, bool) {
	_, t, found := FindIndexIf(items, util.TestOnSecondArg[int, T](predicate))
	return t, found
}

func FindIndexIf[T any, L ~[]T](items L, predicate util.BiPredicate[int, T]) (int, T, bool) {
	for index, item := range items {
		if predicate(index, item) {
			return index, item, true
		}
	}
	var none T
	return -1, none, false
}

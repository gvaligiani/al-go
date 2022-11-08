package list

import "github.com/gvaligiani/al.go/util"

func AnyOf[T any, L ~[]T](items L, predicate util.Predicate[T]) bool {
	_, found := FindIf(items, predicate)
	return found
}

func AnyIndexOf[T any, L ~[]T](items L, predicate util.BiPredicate[int, T]) bool {
	_, _, found := FindIndexIf(items, predicate)
	return found
}

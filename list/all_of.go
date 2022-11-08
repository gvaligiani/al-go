package list

import "github.com/gvaligiani/al.go/util"

func AllOf[T any, L ~[]T](items L, predicate util.Predicate[T]) bool {
	_, found := FindIfNot(items, predicate)
	return !found
}

func AllIndexOf[T any, L ~[]T](items L, predicate util.BiPredicate[int, T]) bool {
	_, _, found := FindIndexIfNot(items, predicate)
	return !found
}

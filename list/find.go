package list

import (
	"github.com/gvaligiani/al.go/util"
)

func Find[T comparable, L ~[]T](items L, value T) bool {
	return FindFn(items, value, util.Equal[T])
}

func DeepFind[T any, L ~[]T](items L, value T) bool {
	return FindFn(items, value, util.DeepEqual[T])
}

func FindFn[T any, L ~[]T](items L, value T, equal util.BiPredicate[T, T]) bool {
	_, found := FindIf(items, func(item T) bool { return equal(item, value) })
	return found
}

func FindIndexOf[T comparable, L ~[]T](items L, value T) (int, bool) {
	return FindFnIndexOf(items, value, util.Equal[T])
}

func DeepFindIndexOf[T any, L ~[]T](items L, value T) (int, bool) {
	return FindFnIndexOf(items, value, util.DeepEqual[T])
}

func FindFnIndexOf[T any, L ~[]T](items L, value T, equal util.BiPredicate[T, T]) (int, bool) {
	index, _, found := FindIndexIf(items, func(_ int, item T) bool { return equal(item, value) })
	return index, found
}

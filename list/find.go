package list

import (
	"github.com/gvaligiani/al.go/util"
)

func Find[T comparable, L ~[]T](items L, value T) (int, bool) {
	return FindFn(items, value, util.Equal[T])
}

func DeepFind[T any, L ~[]T](items L, value T) (int, bool) {
	return FindFn(items, value, util.DeepEqual[T])
}

func FindFn[T any, L ~[]T](items L, value T, equal util.BiPredicate[T, T]) (int, bool) {
	index, _, found := FindIf(items, func(_ int, item T) bool { return equal(item, value) })
	return index, found
}

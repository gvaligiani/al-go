package dict

import (
	"github.com/gvaligiani/al.go/util"
)

func FindValue[K comparable, T comparable, M ~map[K]T](items M, value T) bool {
	return FindValueFn(items, value, util.Equal[T])
}

func DeepFindValue[K comparable, T any, M ~map[K]T](items M, value T) bool {
	return FindValueFn(items, value, util.DeepEqual[T])
}

func FindValueFn[K comparable, T any, M ~map[K]T](items M, value T, equal util.BiPredicate[T, T]) bool {
	_, _, found := FindIf(items, func(_ K, item T) bool { return equal(item, value) })
	return found
}

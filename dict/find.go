package dict

import (
	"github.com/gvaligiani/al.go/util"
)

func Find[K comparable, T comparable, M ~map[K]T](items M, value T) bool {
	return FindFn(items, value, util.Equal[T])
}

func DeepFind[K comparable, T any, M ~map[K]T](items M, value T) bool {
	return FindFn(items, value, util.DeepEqual[T])
}

func FindFn[K comparable, T any, M ~map[K]T](items M, value T, equal util.BiPredicate[T, T]) bool {
	_, found := FindIf(items, func(item T) bool { return equal(item, value) })
	return found
}

func FindKeyOf[K comparable, T comparable, M ~map[K]T](items M, value T) (K, bool) {
	return FindKeyOfFn(items, value, util.Equal[T])
}

func DeepFindKeyOf[K comparable, T any, M ~map[K]T](items M, value T) (K, bool) {
	return FindKeyOfFn(items, value, util.DeepEqual[T])
}

func FindKeyOfFn[K comparable, T any, M ~map[K]T](items M, value T, equal util.BiPredicate[T, T]) (K, bool) {
	key, _, found := FindKeyIf(items, func(key K, item T) bool { return equal(item, value) })
	return key, found
}

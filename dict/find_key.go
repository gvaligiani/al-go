package dict

import (
	"github.com/gvaligiani/al.go/util"
)

func FindKey[K comparable, T any, M ~map[K]T](items M, key K) bool {
	_, found := items[key]
	return found
}

func DeepFindKey[K comparable, T any, M ~map[K]T](items M, key K) bool {
	return FindKeyFn(items, key, util.DeepEqual[K])
}

func FindKeyFn[K comparable, T any, M ~map[K]T](items M, key K, equal util.BiPredicate[K, K]) bool {
	_, _, found := FindIf(items, func(k K, _ T) bool { return equal(k, key) })
	return found
}

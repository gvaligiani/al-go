package dict

import (
	"github.com/gvaligiani/al.go/util"
)

func FindKey[K comparable, V any, D ~map[K]V](d D, k K) bool {
	_, found := d[k]
	return found
}

func DeepFindKey[K comparable, V any, D ~map[K]V](d D, k K) bool {
	return FindKeyFn(d, k, util.DeepEqual[K])
}

func FindKeyFn[K comparable, V any, D ~map[K]V](d D, key K, equal util.BiPredicate[K, K]) bool {
	_, _, found := FindIfKey(d, func(k K, _ V) bool { return equal(k, key) })
	return found
}

func FindValueFromKey[K comparable, V any, D ~map[K]V](d D, key K) (V, bool) {
	value, found := d[key]
	return value, found
}

func DeepFindValueFromKey[K comparable, V any, D ~map[K]V](d D, key K) (V, bool) {
	return FindValueFromKeyFn(d, key, util.DeepEqual[K])
}

func FindValueFromKeyFn[K comparable, V any, D ~map[K]V](d D, key K, equal util.BiPredicate[K, K]) (V, bool) {
	_, value, found := FindIfKey(d, func(k K, _ V) bool { return equal(k, key) })
	return value, found
}

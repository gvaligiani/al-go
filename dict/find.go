package dict

import (
	"github.com/gvaligiani/al.go/util"
)

func Find[K comparable, V comparable, D ~map[K]V](d D, value V) bool {
	return FindFn(d, value, util.Equal[V])
}

func DeepFind[K comparable, V any, D ~map[K]V](d D, value V) bool {
	return FindFn(d, value, util.DeepEqual[V])
}

func FindFn[K comparable, V any, D ~map[K]V](d D, value V, equal util.BiPredicate[V, V]) bool {
	_, found := FindIf(d, func(v V) bool { return equal(v, value) })
	return found
}

func FindKeyFromValue[K comparable, V comparable, D ~map[K]V](d D, value V) (K, bool) {
	return FindKeyFromValueFn(d, value, util.Equal[V])
}

func DeepFindKeyFromValue[K comparable, V any, D ~map[K]V](d D, value V) (K, bool) {
	return FindKeyFromValueFn(d, value, util.DeepEqual[V])
}

func FindKeyFromValueFn[K comparable, V any, D ~map[K]V](d D, value V, equal util.BiPredicate[V, V]) (K, bool) {
	key, _, found := FindIfKey(d, func(_ K, v V) bool { return equal(v, value) })
	return key, found
}

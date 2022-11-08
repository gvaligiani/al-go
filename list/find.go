package list

import (
	"github.com/gvaligiani/al.go/util"
)

func Find[V comparable, L ~[]V](l L, value V) bool {
	return FindFn(l, value, util.Equal[V])
}

func DeepFind[V any, L ~[]V](l L, value V) bool {
	return FindFn(l, value, util.DeepEqual[V])
}

func FindFn[V any, L ~[]V](l L, value V, equal util.BiPredicate[V, V]) bool {
	_, found := FindIf(l, func(v V) bool { return equal(v, value) })
	return found
}

func FindIndexFromValue[V comparable, L ~[]V](l L, value V) (int, bool) {
	return FindIndexFromValueFn(l, value, util.Equal[V])
}

func DeepFindIndexFromValue[V any, L ~[]V](l L, value V) (int, bool) {
	return FindIndexFromValueFn(l, value, util.DeepEqual[V])
}

func FindIndexFromValueFn[V any, L ~[]V](l L, value V, equal util.BiPredicate[V, V]) (int, bool) {
	index, _, found := FindIfIndex(l, func(_ int, v V) bool { return equal(v, value) })
	return index, found
}

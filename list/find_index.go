package list

import (
	"github.com/gvaligiani/al.go/util"
)

func FindIndex[V any, L ~[]V](l L, index int) bool {
	if 0 <= index && index < len(l) {
		return true
	}
	return false
}

func DeepFindIndex[V any, L ~[]V](l L, index int) bool {
	return FindIndexFn(l, index, util.DeepEqual[int])
}

func FindIndexFn[V any, L ~[]V](l L, index int, equal util.BiPredicate[int, int]) bool {
	_, _, found := FindIfIndex(l, func(i int, _ V) bool { return equal(i, index) })
	return found
}

func FindValueFromIndex[V any, L ~[]V](l L, index int) (V, bool) {
	if 0 <= index && index < len(l) {
		return l[index], true
	}
	var noValue V
	return noValue, false
}

func DeepFindValueFromIndex[V any, L ~[]V](l L, index int) (V, bool) {
	return FindValueFromIndexFn(l, index, util.DeepEqual[int])
}

func FindValueFromIndexFn[V any, L ~[]V](l L, index int, equal util.BiPredicate[int, int]) (V, bool) {
	_, value, found := FindIfIndex(l, func(i int, _ V) bool { return equal(i, index) })
	return value, found
}

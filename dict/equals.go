package dict

import (
	"github.com/gvaligiani/al.go/util"
)

func Equal[K comparable, T comparable, M ~map[K]T](left M, right M) bool {
	return EqualFn(left, right, util.Equal[T])
}

func DeepEqual[K comparable, T any, M ~map[K]T](left M, right M) bool {
	return EqualFn(left, right, util.DeepEqual[T])
}

func EqualFn[K comparable, T any, M ~map[K]T](left M, right M, equal util.BiPredicate[T, T]) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for leftKey, leftItem := range left {
		rightItem, found := right[leftKey]
		if !found || !equal(leftItem, rightItem) {
			return false
		}
	}
	return true
}

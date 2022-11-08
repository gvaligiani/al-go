package list

import (
	"github.com/gvaligiani/al.go/util"
)

func Equal[T comparable, L ~[]T](left L, right L) bool {
	return EqualFn(left, right, util.Equal[T])
}

func DeepEqual[T any, L ~[]T](left L, right L) bool {
	return EqualFn(left, right, util.DeepEqual[T])
}

func EqualFn[T any, L ~[]T](left L, right L, equal util.BiPredicate[T, T]) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for leftIndex, leftItem := range left {
		rightItem := right[leftIndex]
		if !equal(leftItem, rightItem) {
			return false
		}
	}
	return true
}

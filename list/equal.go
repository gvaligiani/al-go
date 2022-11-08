package list

import (
	"github.com/gvaligiani/al.go/util"
)

func Equal[V comparable, L ~[]V](left L, right L) bool {
	return EqualFn(left, right, util.Equal[V])
}

func DeepEqual[V any, L ~[]V](left L, right L) bool {
	return EqualFn(left, right, util.DeepEqual[V])
}

func EqualFn[V any, L ~[]V](left L, right L, equal util.BiPredicate[V, V]) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for i, leftValue := range left {
		rightValue := right[i]
		if !equal(leftValue, rightValue) {
			return false
		}
	}
	return true
}

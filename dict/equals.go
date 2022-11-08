package dict

import (
	"github.com/gvaligiani/al.go/util"
)

func Equal[K comparable, V comparable, D ~map[K]V](left D, right D) bool {
	return EqualFn(left, right, util.Equal[V])
}

func DeepEqual[K comparable, V any, D ~map[K]V](left D, right D) bool {
	return EqualFn(left, right, util.DeepEqual[V])
}

func EqualFn[K comparable, V any, D ~map[K]V](left D, right D, equal util.BiPredicate[V, V]) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for k, leftValue := range left {
		rightValue, found := right[k]
		if !found || !equal(leftValue, rightValue) {
			return false
		}
	}
	return true
}

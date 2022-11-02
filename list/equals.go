package list

import (
	"reflect"
)

func Equals[T any, L ~[]T](left L, right L) bool {
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
		if !reflect.DeepEqual(leftItem, rightItem) {
			return false
		}
	}
	return true
}

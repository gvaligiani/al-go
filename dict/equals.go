package dict

import "reflect"

func Equals[K comparable, T any, M ~map[K]T](left M, right M) bool {
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
		if !found || !reflect.DeepEqual(leftItem, rightItem) {
			return false
		}
	}
	return true
}

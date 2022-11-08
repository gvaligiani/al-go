package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func Equal[V comparable, S ~map[V]struct{}](left S, right S) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for leftItem := range left {
		// find by key equality
		if _, found := right[leftItem]; !found {
			return false
		}
	}
	return true
}

func DeepEqual[V comparable, S ~map[V]struct{}](left S, right S) bool {
	return EqualFn(left, right, util.DeepEqual[V])
}

func EqualFn[V comparable, S ~map[V]struct{}](left S, right S, equal util.BiPredicate[V, V]) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for leftItem := range left {
		if !dict.FindKeyFn(right, leftItem, equal) {
			return false
		}
	}
	return true
}

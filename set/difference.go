package set

import (
	"github.com/gvaligiani/al.go/util"
)

func Difference[V comparable, S ~map[V]struct{}](left S, right S) S {
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return Copy(right)
	}
	if right == nil {
		return Copy(left)
	}
	difference := make(S, len(left)+len(right))
	for v := range right {
		if _, found := left[v]; !found {
			difference[v] = struct{}{}
		}
	}
	for v := range left {
		if _, found := right[v]; !found {
			difference[v] = struct{}{}
		}
	}
	return difference
}

func DeepDifference[V comparable, S ~map[V]struct{}](left S, right S) S {
	return DifferenceFn(left, right, util.DeepEqual[V])
}

func DifferenceFn[V comparable, S ~map[V]struct{}](left S, right S, equal util.BiPredicate[V, V]) S {
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return Copy(right)
	}
	if right == nil {
		return Copy(left)
	}
	difference := make(S, len(left)+len(right))
	for v := range right {
		if !FindFn(left, v, equal) {
			difference[v] = struct{}{}
		}
	}
	for v := range left {
		if !FindFn(right, v, equal) {
			difference[v] = struct{}{}
		}
	}
	return difference
}

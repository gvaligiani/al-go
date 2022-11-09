package set

import (
	"github.com/gvaligiani/al.go/util"
)

func Intersection[V comparable, S ~map[V]struct{}](left S, right S) S {
	if left == nil && right == nil {
		return nil
	}
	if left == nil || right == nil {
		return S{}
	}
	intersection := make(S, util.Min(len(left), len(right)))
	for v := range right {
		if _, found := left[v]; found {
			intersection[v] = struct{}{}
		}
	}
	return intersection
}

func DeepIntersection[V comparable, S ~map[V]struct{}](left S, right S) S {
	return IntersectionFn(left, right, util.DeepEqual[V])
}

func IntersectionFn[V comparable, S ~map[V]struct{}](left S, right S, equal util.BiPredicate[V, V]) S {
	if left == nil && right == nil {
		return nil
	}
	if left == nil || right == nil {
		return S{}
	}
	intersection := make(S, util.Min(len(left), len(right)))
	for v := range right {
		if FindFn(left, v, equal) {
			intersection[v] = struct{}{}
		}
	}
	return intersection
}

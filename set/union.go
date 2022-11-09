package set

import (
	"github.com/gvaligiani/al.go/util"
)

func Union[V comparable, S ~map[V]struct{}](left S, right S) S {
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return Copy(right)
	}
	union := Copy(left)
	for v := range right {
		union[v] = struct{}{}
	}
	return union
}

func DeepUnion[V comparable, S ~map[V]struct{}](left S, right S) S {
	return UnionFn(left, right, util.DeepEqual[V])
}

func UnionFn[V comparable, S ~map[V]struct{}](left S, right S, equal util.BiPredicate[V, V]) S {
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return Copy(right)
	}
	union := Copy(left)
	for v := range right {
		if !FindFn(left, v, equal) {
			union[v] = struct{}{}
		}
	}
	return union
}

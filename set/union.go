package set

import (
	"github.com/gvaligiani/al.go/util"
)

func Union[V comparable, S ~map[V]struct{}](left S, right S) S {
	if left == nil {
		return Copy(right)
	}
	union := Copy(left)
	Merge(Ptr(union), right)
	return union
}

func DeepUnion[V comparable, S ~map[V]struct{}](left S, right S) S {
	return UnionFn(left, right, util.DeepEqual[V])
}

func UnionFn[V comparable, S ~map[V]struct{}](left S, right S, equal util.BiPredicate[V, V]) S {
	if left == nil {
		return Copy(right)
	}
	union := Copy(left)
	MergeFn(Ptr(union), right, equal)
	return union
}

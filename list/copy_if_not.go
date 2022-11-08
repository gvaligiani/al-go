package list

import "github.com/gvaligiani/al.go/util"

func CopyIfNot[V any, L ~[]V](l L, predicate util.Predicate[V]) L {
	return CopyIf(l, util.Not(predicate))
}

func CopyIfNotIndex[V any, L ~[]V](l L, predicate util.BiPredicate[int, V]) L {
	return CopyIfIndex(l, util.BiNot(predicate))
}

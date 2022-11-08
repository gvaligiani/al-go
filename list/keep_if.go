package list

import "github.com/gvaligiani/al.go/util"

func KeepIf[V any, L ~[]V](l *L, predicate util.Predicate[V]) bool {
	return RemoveIf(l, util.Not(predicate))
}

func KeepIfIndex[V any, L ~[]V](l *L, predicate util.BiPredicate[int, V]) bool {
	return RemoveIfIndex(l, util.BiNot(predicate))
}

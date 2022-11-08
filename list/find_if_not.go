package list

import "github.com/gvaligiani/al.go/util"

func FindIfNot[V any, L ~[]V](l L, predicate util.Predicate[V]) (V, bool) {
	return FindIf(l, util.Not(predicate))
}

func FindIfNotIndex[V any, L ~[]V](l L, predicate util.BiPredicate[int, V]) (int, V, bool) {
	return FindIfIndex(l, util.BiNot(predicate))
}

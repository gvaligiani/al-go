package list

import "github.com/gvaligiani/al.go/util"

func AllOf[V any, L ~[]V](l L, predicate util.Predicate[V]) bool {
	_, found := FindIfNot(l, predicate)
	return !found
}

func AllIndexOf[V any, L ~[]V](l L, predicate util.BiPredicate[int, V]) bool {
	_, _, found := FindIfNotIndex(l, predicate)
	return !found
}

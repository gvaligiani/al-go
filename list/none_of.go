package list

import "github.com/gvaligiani/al.go/util"

func NoneOf[V any, L ~[]V](l L, predicate util.Predicate[V]) bool {
	_, found := FindIf(l, predicate)
	return !found
}

func NoIndexOf[V any, L ~[]V](l L, predicate util.BiPredicate[int, V]) bool {
	_, _, found := FindIfIndex(l, predicate)
	return !found
}

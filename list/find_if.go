package list

import "github.com/gvaligiani/al.go/util"

func FindIf[V any, L ~[]V](l L, predicate util.Predicate[V]) (V, bool) {
	_, t, found := FindIfIndex(l, util.TestOnSecondArg[int](predicate))
	return t, found
}

func FindIfIndex[V any, L ~[]V](l L, predicate util.BiPredicate[int, V]) (int, V, bool) {
	for i, v := range l {
		if predicate(i, v) {
			return i, v, true
		}
	}
	var none V
	return -1, none, false
}

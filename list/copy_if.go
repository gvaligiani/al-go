package list

import "github.com/gvaligiani/al.go/util"

func CopyIf[V any, L ~[]V](l L, predicate util.Predicate[V]) L {
	return CopyIfIndex(l, util.TestOnSecondArg[int](predicate))
}

func CopyIfIndex[V any, L ~[]V](l L, predicate util.BiPredicate[int, V]) L {
	if l == nil {
		return nil
	}
	copy := make(L, 0, len(l))
	for i, t := range l {
		if predicate(i, t) {
			copy = append(copy, t)
		}
	}
	return copy
}

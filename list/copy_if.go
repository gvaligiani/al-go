package list

import "github.com/gvaligiani/al.go/util"

func CopyIf[T any, L ~[]T](l L, predicate util.Predicate[T]) L {
	return CopyIndexIf(l, util.TestOnSecondArg[int, T](predicate))
}

func CopyIndexIf[T any, L ~[]T](l L, predicate util.BiPredicate[int, T]) L {
	if l == nil {
		return nil
	}
	result := make(L, 0, len(l))
	for index, t := range l {
		if predicate(index, t) {
			result = append(result, t)
		}
	}
	return result
}

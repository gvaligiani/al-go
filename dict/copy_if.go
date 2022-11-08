package dict

import "github.com/gvaligiani/al.go/util"

func CopyIf[K comparable, T any, M ~map[K]T](items M, predicate util.Predicate[T]) M {
	return CopyKeyIf(items, util.TestOnSecondArg[K, T](predicate))
}

func CopyKeyIf[K comparable, T any, M ~map[K]T](items M, predicate util.BiPredicate[K, T]) M {
	if items == nil {
		return nil
	}
	result := make(M, len(items))
	for key, item := range items {
		if predicate(key, item) {
			result[key] = item
		}
	}
	return result
}

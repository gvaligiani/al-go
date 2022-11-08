package dict

import "github.com/gvaligiani/al.go/util"

func FindIf[K comparable, T any, M ~map[K]T](items M, predicate util.Predicate[T]) (T, bool) {
	_, t, found := FindKeyIf(items, util.TestOnSecondArg[K, T](predicate))
	return t, found
}

func FindKeyIf[K comparable, T any, M ~map[K]T](items M, predicate util.BiPredicate[K, T]) (K, T, bool) {
	for key, item := range items {
		if predicate(key, item) {
			return key, item, true
		}
	}
	var noKey K
	var noValue T
	return noKey, noValue, false
}

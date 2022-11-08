package dict

import "github.com/gvaligiani/al.go/util"

func FindIfNot[K comparable, T any, M ~map[K]T](items M, predicate util.Predicate[T]) (T, bool) {
	return FindIf(items, util.Not(predicate))
}

func FindKeyIfNot[K comparable, T any, M ~map[K]T](items M, predicate util.BiPredicate[K, T]) (K, T, bool) {
	return FindKeyIf(items, util.BiNot(predicate))
}

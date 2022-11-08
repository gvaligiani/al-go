package dict

import "github.com/gvaligiani/al.go/util"

func CopyIfNot[K comparable, T any, M ~map[K]T](m M, predicate util.Predicate[T]) M {
	return CopyIf(m, util.Not(predicate))
}

func CopyKeyIfNot[K comparable, T any, M ~map[K]T](m M, predicate util.BiPredicate[K, T]) M {
	return CopyKeyIf(m, util.BiNot(predicate))
}

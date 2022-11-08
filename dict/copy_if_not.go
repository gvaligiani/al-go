package dict

import "github.com/gvaligiani/al.go/util"

func CopyIfNot[K comparable, V any, D ~map[K]V](d D, predicate util.Predicate[V]) D {
	return CopyIf(d, util.Not(predicate))
}

func CopyIfNotKey[K comparable, V any, D ~map[K]V](d D, predicate util.BiPredicate[K, V]) D {
	return CopyIfKey(d, util.BiNot(predicate))
}

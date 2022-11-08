package dict

import "github.com/gvaligiani/al.go/util"

func FindIfNot[K comparable, V any, D ~map[K]V](d D, predicate util.Predicate[V]) (V, bool) {
	return FindIf(d, util.Not(predicate))
}

func FindIfNotKey[K comparable, V any, D ~map[K]V](d D, predicate util.BiPredicate[K, V]) (K, V, bool) {
	return FindIfKey(d, util.BiNot(predicate))
}

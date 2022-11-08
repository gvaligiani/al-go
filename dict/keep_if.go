package dict

import "github.com/gvaligiani/al.go/util"

func KeepIf[K comparable, V any, D ~map[K]V](d *D, predicate util.Predicate[V]) bool {
	return RemoveIf(d, util.Not(predicate))
}

func KeepIfKey[K comparable, V any, D ~map[K]V](d *D, predicate util.BiPredicate[K, V]) bool {
	return RemoveIfKey(d, util.BiNot(predicate))
}

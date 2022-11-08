package dict

import "github.com/gvaligiani/al.go/util"

func AllOf[K comparable, V any, D ~map[K]V](d D, predicate util.Predicate[V]) bool {
	_, found := FindIfNot(d, predicate)
	return !found
}

func AllKeyOf[K comparable, V any, D ~map[K]V](d D, predicate util.BiPredicate[K, V]) bool {
	_, _, found := FindIfNotKey(d, predicate)
	return !found
}

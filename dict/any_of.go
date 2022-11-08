package dict

import "github.com/gvaligiani/al.go/util"

func AnyOf[K comparable, V any, D ~map[K]V](d D, predicate util.Predicate[V]) bool {
	_, found := FindIf(d, predicate)
	return found
}

func AnyKeyOf[K comparable, V any, D ~map[K]V](d D, predicate util.BiPredicate[K, V]) bool {
	_, _, found := FindIfKey(d, predicate)
	return found
}

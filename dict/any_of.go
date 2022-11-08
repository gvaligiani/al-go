package dict

import "github.com/gvaligiani/al.go/util"

func AnyOf[K comparable, T any, M ~map[K]T](items M, predicate util.Predicate[T]) bool {
	_, found := FindIf(items, predicate)
	return found
}

func AnyKeyOf[K comparable, T any, M ~map[K]T](items M, predicate util.BiPredicate[K, T]) bool {
	_, _, found := FindKeyIf(items, predicate)
	return found
}

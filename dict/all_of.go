package dict

import "github.com/gvaligiani/al.go/util"

func AllOf[K comparable, T any, M ~map[K]T](items M, predicate util.Predicate[T]) bool {
	_, found := FindIfNot(items, predicate)
	return !found
}

func AllKeyOf[K comparable, T any, M ~map[K]T](items M, predicate util.BiPredicate[K, T]) bool {
	_, _, found := FindKeyIfNot(items, predicate)
	return !found
}

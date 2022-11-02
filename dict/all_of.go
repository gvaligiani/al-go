package dict

func AllOf[K comparable, T any, M ~map[K]T](items M, predicate Predicate[K, T]) bool {
	_, _, found := FindIfNot(items, predicate)
	return !found
}

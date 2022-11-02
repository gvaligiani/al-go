package dict

func NoneOf[K comparable, T any, M ~map[K]T](items M, predicate Predicate[K, T]) bool {
	_, _, found := FindIf(items, predicate)
	return !found
}

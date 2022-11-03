package list

func AllOf[T any, L ~[]T](items L, predicate Predicate[T]) bool {
	_, found := FindIfNot(items, predicate)
	return !found
}

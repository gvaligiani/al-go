package list

func AllOf[T any, L ~[]T](items L, predicate Predicate[T]) bool {
	_, _, found := FindIfNot(items, predicate)
	return !found
}

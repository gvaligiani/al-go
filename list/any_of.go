package list

func AnyOf[T any, L ~[]T](items L, predicate Predicate[T]) bool {
	_, _, found := FindIf(items, predicate)
	return found
}

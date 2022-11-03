package list

func FindIfNot[T any, L ~[]T](items L, predicate Predicate[T]) (T, bool) {
	return FindIf(items, Not(predicate))
}

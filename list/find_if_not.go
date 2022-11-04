package list

func FindIfNot[T any, L ~[]T](items L, predicate Predicate[T]) (int, T, bool) {
	return FindIf(items, Not(predicate))
}

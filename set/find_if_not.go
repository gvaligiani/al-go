package set

func FindIfNot[T comparable, S ~map[T]struct{}](items S, predicate Predicate[T]) (T, bool) {
	return FindIf(items, Not(predicate))
}

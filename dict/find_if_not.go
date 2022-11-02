package dict

func FindIfNot[K comparable, T any, M ~map[K]T](items M, predicate Predicate[K, T]) (K, T, bool) {
	return FindIf(items, Not(predicate))
}

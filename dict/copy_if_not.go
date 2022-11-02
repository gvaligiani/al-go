package dict

func CopyIfNot[K comparable, T any, M ~map[K]T](items M, predicate Predicate[K, T]) M {
	return CopyIf(items, Not(predicate))
}

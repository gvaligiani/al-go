package dict

func KeepIf[K comparable, T any, M ~map[K]T](items *M, predicate Predicate[K, T]) bool {
	return RemoveIf(items, Not(predicate))
}

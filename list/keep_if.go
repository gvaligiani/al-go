package list

func KeepIf[T any, L ~[]T](items *L, predicate Predicate[T]) bool {
	return RemoveIf(items, Not(predicate))
}

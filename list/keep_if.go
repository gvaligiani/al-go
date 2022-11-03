package list

func KeepIf[T any, L ~[]T](items *L, predicate Predicate[T]) {
	RemoveIf(items, Not(predicate))
}

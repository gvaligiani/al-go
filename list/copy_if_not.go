package list

func CopyIfNot[T any, L ~[]T](items L, predicate Predicate[T]) L {
	return CopyIf(items, Not(predicate))
}

package set

func CopyIfNot[T comparable, S ~map[T]struct{}](items S, predicate Predicate[T]) S {
	return CopyIf(items, Not(predicate))
}

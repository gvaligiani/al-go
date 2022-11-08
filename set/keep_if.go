package set

func KeepIf[T comparable, S ~map[T]struct{}](items *S, predicate Predicate[T]) bool {
	return RemoveIf(items, func(value T) bool { return !predicate(value) })
}

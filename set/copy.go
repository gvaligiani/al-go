package set

func Copy[T comparable, S ~map[T]struct{}](items S) S {
	return CopyIf(items, True[T]())
}

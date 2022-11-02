package list

func Copy[T any, L ~[]T](items L) L {
	return CopyIf(items, True[T]())
}

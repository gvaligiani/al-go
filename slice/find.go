package slice

func Find[T comparable](items []T, value T) (T, bool) {
	return FindIf(items, func(item T) bool { return item == value })
}
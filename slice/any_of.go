package slice

func AnyOf[T any](items []T, predicate func(T) bool) bool {
	_, found := FindIf(items, predicate)
	return found
}
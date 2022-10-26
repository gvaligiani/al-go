package slice

func NoneOf[T any](items []T, predicate func(T) bool) bool {
	_, found := FindIf(items, predicate)
	return !found
}
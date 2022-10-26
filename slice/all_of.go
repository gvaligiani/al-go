package slice

func AllOf[T any](items []T, predicate func(T) bool) bool {
	_, found := FindIfNot(items, predicate)
	return !found
}
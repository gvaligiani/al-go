package list

func FindIf[T any](items []T, predicate func(T) bool) (T, bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	var none T
	return none, false
}
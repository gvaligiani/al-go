package list

func FindIf[T any, L ~[]T](items L, predicate Predicate[T]) (T, bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	var none T
	return none, false
}
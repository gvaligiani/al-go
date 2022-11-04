package list

func FindIf[T any, L ~[]T](items L, predicate Predicate[T]) (T, bool) {
	for index, item := range items {
		if predicate(index, item) {
			return item, true
		}
	}
	var none T
	return none, false
}

package list

func FindIf[T any, L ~[]T](items L, predicate Predicate[T]) (int, T, bool) {
	for index, item := range items {
		if predicate(index, item) {
			return index, item, true
		}
	}
	var none T
	return -1, none, false
}

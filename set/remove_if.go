package set

func RemoveIf[T comparable, S ~map[T]struct{}](items *S, predicate Predicate[T]) {
	if len(*items) == 0 {
		return
	}
	itemsToRemove := make([]T, 0, len(*items))
	for item := range *items {
		if predicate(item) {
			itemsToRemove = append(itemsToRemove, item)
		}
	}
	for _, itemToRemove := range itemsToRemove {
		delete(*items, itemToRemove)
	}
}

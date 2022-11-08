package set

func RemoveIf[T comparable, S ~map[T]struct{}](items *S, predicate Predicate[T]) bool {
	if len(*items) == 0 {
		return false
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
	return len(itemsToRemove) > 0
}

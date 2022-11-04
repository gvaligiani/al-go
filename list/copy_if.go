package list

func CopyIf[T any, L ~[]T](items L, predicate Predicate[T]) L {
	if items == nil {
		return nil
	}
	result := make(L, 0, len(items))
	for index, item := range items {
		if predicate(index, item) {
			result = append(result, item)
		}
	}
	return result
}

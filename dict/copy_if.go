package dict

func CopyIf[K comparable, T any, M ~map[K]T](items M, predicate Predicate[K, T]) M {
	if items == nil {
		return nil
	}
	result := make(M, len(items))
	for key, item := range items {
		if predicate(key, item) {
			result[key] = item
		}
	}
	return result
}

package set

func CopyIf[T comparable, S ~map[T]struct{}](items S, predicate Predicate[T]) S {
	if items == nil {
		return nil
	}
	result := make(S, len(items))
	for item := range items {
		if predicate(item) {
			result[item] = struct{}{}
		}
	}
	return result
}

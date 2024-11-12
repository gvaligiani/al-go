package set

// //////////////////////////////////////////////////
// find if

func FindIf[T comparable](items map[T]struct{}, predicate func(item T) bool) (T, bool) {
	for item := range items {
		if predicate(item) {
			return item, true
		}
	}
	var empty T
	return empty, false
}

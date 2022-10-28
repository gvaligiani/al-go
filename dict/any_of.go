package dict

func AnyOf[K comparable,T any](items map[K]T, predicate func(K,T) bool) bool {
	_, _, found := FindIf(items, predicate)
	return found
}
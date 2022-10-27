package dict

func AllOf[K comparable,T any](items map[K]T, predicate func(K,T) bool) bool {
	_, found := FindIfNot(items, predicate)
	return !found
}
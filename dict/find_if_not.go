package dict

func FindIfNot[K comparable,T any](items map[K]T, predicate func(K,T) bool) (T, bool) {
	for key, item := range items {
		if !predicate(key,item) {
			return item, true
		}
	}
	var none T
	return none, false
}
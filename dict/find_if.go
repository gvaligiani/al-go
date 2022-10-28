package dict

func FindIf[K comparable,T any](items map[K]T, predicate func(K,T) bool) (K, T, bool) {
	for key, item := range items {
		if predicate(key,item) {
			return key, item, true
		}
	}
	var noKey K
	var noValue T
	return noKey, noValue, false
}
package dict

func FindIf[K comparable, T any, M ~map[K]T](items M, predicate Predicate[K, T]) (K, T, bool) {
	for key, item := range items {
		if predicate(key, item) {
			return key, item, true
		}
	}
	var noKey K
	var noValue T
	return noKey, noValue, false
}

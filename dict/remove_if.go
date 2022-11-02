package dict

func RemoveIf[K comparable, T any, M ~map[K]T](items *M, predicate Predicate[K, T]) {
	if len(*items) == 0 {
		return
	}
	keysToRemove := make([]K, 0, len(*items))
	for key, item := range *items {
		if predicate(key, item) {
			keysToRemove = append(keysToRemove, key)
		}
	}
	for _, keyToRemove := range keysToRemove {
		delete(*items, keyToRemove)
	}
}

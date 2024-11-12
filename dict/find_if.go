package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// find if

func FindIf[K comparable, V any](items map[K]V, predicate fn.BiPredicate[K, V]) (K, V, bool) {
	for key, value := range items {
		if predicate(key, value) {
			return key, value, true
		}
	}
	var emptyKey K
	var emptyValue V
	return emptyKey, emptyValue, false
}

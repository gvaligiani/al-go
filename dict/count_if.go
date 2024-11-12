package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// count if

func CountIf[K comparable, V any](items map[K]V, predicate fn.BiPredicate[K, V]) int {
	count := 0
	for key, value := range items {
		if predicate(key, value) {
			count++
		}
	}
	return count
}

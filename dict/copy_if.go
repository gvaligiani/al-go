package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// copy if

func CopyIf[K comparable, V any](items map[K]V, predicate fn.BiPredicate[K, V]) map[K]V {
	copy := make(map[K]V, len(items))
	for key, value := range items {
		if predicate(key, value) {
			copy[key] = value
		}
	}
	return copy
}

package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// remove if

func RemoveIf[K comparable, V any, M ~map[K]V](items *M, predicate fn.BiPredicate[K, V]) {
	if items == nil || len(*items) == 0 {
		return
	}
	for key, value := range *items {
		if predicate(key, value) {
			delete(*items, key)
		}
	}
}

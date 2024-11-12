package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// remove if

func RemoveIf[T comparable, S ~map[T]struct{}](items *S, predicate fn.Predicate[T]) {
	if items == nil || len(*items) == 0 {
		return
	}
	for key := range *items {
		if predicate(key) {
			delete(*items, key)
		}
	}
}

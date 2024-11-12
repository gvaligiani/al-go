package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// copy if

func CopyIf[T comparable](items map[T]struct{}, predicate fn.Predicate[T]) map[T]struct{} {
	copy := make(map[T]struct{}, len(items))
	for item := range items {
		if predicate(item) {
			copy[item] = struct{}{}
		}
	}
	return copy
}

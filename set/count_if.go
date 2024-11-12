package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// count if

func CountIf[T comparable](items map[T]struct{}, predicate fn.Predicate[T]) int {
	count := 0
	for item := range items {
		if predicate(item) {
			count++
		}
	}
	return count
}

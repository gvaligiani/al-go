package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// count if

func CountIf[T any](items []T, predicate fn.Predicate[T]) int {
	count := 0
	for _, item := range items {
		if predicate(item) {
			count++
		}
	}
	return count
}

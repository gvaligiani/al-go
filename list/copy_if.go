package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// copy if

func CopyIf[T any](items []T, predicate fn.Predicate[T]) []T {
	filtered := make([]T, 0, len(items))
	for _, item := range items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

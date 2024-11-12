package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// any

func AnyOf[T any](items []T, predicate fn.Predicate[T]) bool {
	_, found := FindIf(items, predicate)
	return found
}

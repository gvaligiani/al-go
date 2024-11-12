package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// all

func AllOf[T any](items []T, predicate fn.Predicate[T]) bool {
	_, found := FindIf(items, fn.Not(predicate))
	return !found
}

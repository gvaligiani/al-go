package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// all

func AllOf[T comparable](items map[T]struct{}, predicate fn.Predicate[T]) bool {
	_, found := FindIf(items, fn.Not(predicate))
	return !found
}

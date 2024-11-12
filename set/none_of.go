package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// none of

func NoneOf[T comparable](items map[T]struct{}, predicate fn.Predicate[T]) bool {
	_, found := FindIf(items, predicate)
	return !found
}

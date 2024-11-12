package list

import "github.com/gvaligiani/al-go/fn"

func NoneOf[T any](items []T, predicate fn.Predicate[T]) bool {
	_, found := FindIf(items, predicate)
	return !found
}

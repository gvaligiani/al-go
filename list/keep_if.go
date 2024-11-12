package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// keep if

func KeepIf[T any, L ~[]T](items *L, predicate fn.Predicate[T]) {
	RemoveIf(items, fn.Not(predicate))
}

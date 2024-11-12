package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// keep if

func KeepIf[T comparable, S ~map[T]struct{}](items *S, predicate fn.Predicate[T]) {
	RemoveIf(items, fn.Not(predicate))
}

package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// keep if

func KeepIf[K comparable, V any, M ~map[K]V](items *M, predicate fn.BiPredicate[K, V]) {
	RemoveIf(items, fn.BiNot(predicate))
}

package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// all of

func AllOf[K comparable, V any](items map[K]V, predicate fn.BiPredicate[K, V]) bool {
	_, _, found := FindIf(items, fn.BiNot(predicate))
	return !found
}

package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// any of

func AnyOf[K comparable, V any](items map[K]V, predicate fn.BiPredicate[K, V]) bool {
	_, _, found := FindIf(items, predicate)
	return found
}

package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// none of

func NoneOf[K comparable, V any](items map[K]V, predicate fn.BiPredicate[K, V]) bool {
	_, _, found := FindIf(items, predicate)
	return !found
}

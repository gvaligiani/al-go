package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// reduce

func Reduce[K comparable, V any, U any](items map[K]V, initial U, reduce fn.BiReducer[K, V, U]) U {
	next := initial
	for key, value := range items {
		next = reduce(key, value, next)
	}
	return next
}

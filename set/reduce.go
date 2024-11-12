package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// reduce

func Reduce[T comparable, U any](items map[T]struct{}, initial U, reduce fn.Reducer[T, U]) U {
	next := initial
	for item := range items {
		next = reduce(item, next)
	}
	return next
}

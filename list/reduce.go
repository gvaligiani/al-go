package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// reduce

func Reduce[T any, U any](items []T, initial U, reduce fn.Reducer[T, U]) U {
	next := initial
	for _, item := range items {
		next = reduce(item, next)
	}
	return next
}

package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// copy if

func Copy[T comparable](items map[T]struct{}, predicate fn.Predicate[T]) map[T]struct{} {
	copy := make(map[T]struct{}, len(items))
	for item := range items {
		copy[item] = struct{}{}
	}
	return copy
}

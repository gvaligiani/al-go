package set

import "github.com/gvaligiani/al.go/dict"

func AnyOf[T comparable, S ~map[T]struct{}](items S, predicate Predicate[T]) bool {
	return dict.AnyOf(items, func(item T, _ struct{}) bool { return predicate(item) })
}

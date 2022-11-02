package set

import "github.com/gvaligiani/al.go/dict"

func NoneOf[T comparable, S ~map[T]struct{}](items S, predicate Predicate[T]) bool {
	return dict.NoneOf(items, func(item T, _ struct{}) bool { return predicate(item) })
}

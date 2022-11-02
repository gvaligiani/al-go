package set

import "github.com/gvaligiani/al.go/dict"

func FindIf[T comparable, S ~map[T]struct{}](items S, predicate Predicate[T]) (T, bool) {
	item, _, found := dict.FindIf(items, func(item T, _ struct{}) bool { return predicate(item) })
	return item, found
}

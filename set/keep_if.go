package set

import "github.com/gvaligiani/al.go/util"

func KeepIf[T comparable, S ~map[T]struct{}](items *S, predicate util.Predicate[T]) bool {
	return RemoveIf(items, func(value T) bool { return !predicate(value) })
}

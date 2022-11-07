package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func Find[T comparable, S ~map[T]struct{}](items S, value T) bool {
	_, found := items[value]
	return found
}

func DeepFind[T comparable, S ~map[T]struct{}](items S, value T) bool {
	return FindFn(items, value, util.DeepEqual[T])
}

func FindFn[T comparable, S ~map[T]struct{}](items S, value T, equal util.BiPredicate[T, T]) bool {
	return dict.FindKeyFn(items, value, equal)
}

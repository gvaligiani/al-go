package set

import "github.com/gvaligiani/al.go/util"

func Copy[T comparable, S ~map[T]struct{}](items S) S {
	return CopyIf(items, util.True[T]())
}

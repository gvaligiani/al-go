package list

import "github.com/gvaligiani/al.go/util"

func Copy[T any, L ~[]T](items L) L {
	return CopyIf(items, util.True[T]())
}

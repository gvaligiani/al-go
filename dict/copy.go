package dict

import "github.com/gvaligiani/al.go/util"

func Copy[K comparable, T any, M ~map[K]T](items M) M {
	return CopyIf(items, util.True[T]())
}

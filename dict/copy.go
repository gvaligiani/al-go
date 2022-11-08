package dict

import "github.com/gvaligiani/al.go/util"

func Copy[K comparable, V any, D ~map[K]V](d D) D {
	return CopyIf(d, util.True[V]())
}

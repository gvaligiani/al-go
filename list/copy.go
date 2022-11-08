package list

import "github.com/gvaligiani/al.go/util"

func Copy[V any, L ~[]V](l L) L {
	return CopyIf(l, util.True[V]())
}

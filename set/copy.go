package set

import "github.com/gvaligiani/al.go/util"

func Copy[V comparable, S ~map[V]struct{}](s S) S {
	return CopyIf(s, util.True[V]())
}

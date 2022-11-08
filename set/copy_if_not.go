package set

import "github.com/gvaligiani/al.go/util"

func CopyIfNot[V comparable, S ~map[V]struct{}](s S, predicate util.Predicate[V]) S {
	return CopyIf(s, util.Not(predicate))
}

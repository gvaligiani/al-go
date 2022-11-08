package set

import "github.com/gvaligiani/al.go/util"

func KeepIf[V comparable, S ~map[V]struct{}](s *S, predicate util.Predicate[V]) bool {
	return RemoveIf(s, func(v V) bool { return !predicate(v) })
}

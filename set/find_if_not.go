package set

import "github.com/gvaligiani/al.go/util"

func FindIfNot[V comparable, S ~map[V]struct{}](s S, predicate util.Predicate[V]) (V, bool) {
	return FindIf(s, util.Not(predicate))
}

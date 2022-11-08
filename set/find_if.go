package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func FindIf[V comparable, S ~map[V]struct{}](s S, predicate util.Predicate[V]) (V, bool) {
	v, _, found := dict.FindIfKey(s, util.TestOnFirstArg[V, struct{}](predicate))
	return v, found
}

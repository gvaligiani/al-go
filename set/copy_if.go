package set

import "github.com/gvaligiani/al.go/util"

func CopyIf[V comparable, S ~map[V]struct{}](s S, predicate util.Predicate[V]) S {
	if s == nil {
		return nil
	}
	copy := make(S, len(s))
	for v := range s {
		if predicate(v) {
			copy[v] = struct{}{}
		}
	}
	return copy
}

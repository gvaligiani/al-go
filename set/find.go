package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func Find[V comparable, S ~map[V]struct{}](s S, v V) bool {
	_, found := s[v]
	return found
}

func DeepFind[V comparable, S ~map[V]struct{}](s S, v V) bool {
	return FindFn(s, v, util.DeepEqual[V])
}

func FindFn[V comparable, S ~map[V]struct{}](s S, v V, equal util.BiPredicate[V, V]) bool {
	return dict.FindKeyFn(s, v, equal)
}

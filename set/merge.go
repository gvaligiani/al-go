package set

import (
	"github.com/gvaligiani/al.go/util"
)

func Merge[V comparable, S ~map[V]struct{}](target *S, source S) bool {
	if target == nil || len(source) == 0 {
		return false
	}
	updated := false
	for v := range source {
		if _, found := (*target)[v]; !found {
			updated = true
			(*target)[v] = struct{}{}
		}
	}
	return updated
}

func DeepMerge[V comparable, S ~map[V]struct{}](target *S, source S) bool {
	return MergeFn(target, source, util.DeepEqual[V])
}

func MergeFn[V comparable, S ~map[V]struct{}](target *S, source S, equal util.BiPredicate[V, V]) bool {
	if target == nil || len(source) == 0 {
		return false
	}
	updated := false
	for v := range source {
		if !FindFn(*target, v, equal) {
			updated = true
			(*target)[v] = struct{}{}
		}
	}
	return updated
}

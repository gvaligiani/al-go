package set

import "github.com/gvaligiani/al.go/util"

func RemoveIf[V comparable, S ~map[V]struct{}](s *S, predicate util.Predicate[V]) bool {
	if s == nil || len(*s) == 0 {
		return false
	}
	toRemove := make([]V, 0, len(*s))
	for v := range *s {
		if predicate(v) {
			toRemove = append(toRemove, v)
		}
	}
	for _, v := range toRemove {
		delete(*s, v)
	}
	return len(toRemove) > 0
}

package dict

import "github.com/gvaligiani/al.go/util"

func RemoveIf[K comparable, T any, M ~map[K]T](items *M, predicate util.Predicate[T]) bool {
	return RemoveKeyIf(items, func(_ K, t T) bool { return predicate(t) })
}

func RemoveKeyIf[K comparable, T any, M ~map[K]T](m *M, predicate util.BiPredicate[K, T]) bool {
	if len(*m) == 0 {
		return false
	}
	toRemove := make([]K, 0, len(*m))
	for k, t := range *m {
		if predicate(k, t) {
			toRemove = append(toRemove, k)
		}
	}
	for _, k := range toRemove {
		delete(*m, k)
	}
	return len(toRemove) > 0
}

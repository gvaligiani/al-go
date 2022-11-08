package dict

import "github.com/gvaligiani/al.go/util"

func RemoveIf[K comparable, V any, D ~map[K]V](d *D, predicate util.Predicate[V]) bool {
	return RemoveIfKey(d, func(_ K, v V) bool { return predicate(v) })
}

func RemoveIfKey[K comparable, V any, D ~map[K]V](d *D, predicate util.BiPredicate[K, V]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	keys := make([]K, 0, len(*d))
	for k, t := range *d {
		if predicate(k, t) {
			keys = append(keys, k)
		}
	}
	for _, k := range keys {
		delete(*d, k)
	}
	return len(keys) > 0
}

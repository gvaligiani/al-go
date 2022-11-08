package dict

import "github.com/gvaligiani/al.go/util"

func CopyIf[K comparable, V any, D ~map[K]V](d D, predicate util.Predicate[V]) D {
	return CopyIfKey(d, util.TestOnSecondArg[K](predicate))
}

func CopyIfKey[K comparable, V any, D ~map[K]V](d D, predicate util.BiPredicate[K, V]) D {
	if d == nil {
		return nil
	}
	copy := make(D, len(d))
	for k, v := range d {
		if predicate(k, v) {
			copy[k] = v
		}
	}
	return copy
}

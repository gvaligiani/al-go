package dict

import "github.com/gvaligiani/al.go/util"

func FindIf[K comparable, V any, D ~map[K]V](d D, predicate util.Predicate[V]) (V, bool) {
	_, t, found := FindIfKey(d, util.TestOnSecondArg[K](predicate))
	return t, found
}

func FindIfKey[K comparable, V any, D ~map[K]V](d D, predicate util.BiPredicate[K, V]) (K, V, bool) {
	for k, v := range d {
		if predicate(k, v) {
			return k, v, true
		}
	}
	var noKey K
	var noValue V
	return noKey, noValue, false
}

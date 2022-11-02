package dict

import "reflect"

func FindKey[K comparable, T any, M ~map[K]T](items M, value K) bool {
	_, _, found := FindIf(items, func(key K, _ T) bool { return reflect.DeepEqual(key, value) })
	return found
}
